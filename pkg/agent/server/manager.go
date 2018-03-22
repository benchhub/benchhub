package server

import (
	"context"
	igrpc "github.com/at15/go.ice/ice/transport/grpc"
	ihttp "github.com/at15/go.ice/ice/transport/http"
	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"
	"google.golang.org/grpc"
	"sync"

	"github.com/benchhub/benchhub/pkg/agent/config"
	mygrpc "github.com/benchhub/benchhub/pkg/agent/transport/grpc"
	crpc "github.com/benchhub/benchhub/pkg/central/transport/grpc"
)

type Manager struct {
	cfg config.ServerConfig

	registry *Registry

	grpcSrv       *GrpcServer
	grpcTransport *igrpc.Server
	httpSrv       *HttpServer
	httpTransport *ihttp.Server

	client crpc.BenchHubCentralClient
	beater *Beater
	log    *dlog.Logger
}

func NewManager(cfg config.ServerConfig) (*Manager, error) {
	log.Info("creating benchhub agent manager")

	r := NewRegistry(cfg)

	// state machine
	state, err := NewStateMachine()
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create state machine")
	}
	r.State = state

	// client for central
	conn, err := grpc.Dial(cfg.Central.Addr, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "can't dial central server")
	}
	client := crpc.NewClient(conn)

	// beater
	beater, err := NewBeater(client, r)
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create beater")
	}

	// grpc and http server
	grpcSrv, err := NewGrpcServer(r)
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create grpc server")
	}
	grpcTransport, err := igrpc.NewServer(cfg.Grpc, func(s *grpc.Server) {
		mygrpc.RegisterBenchHubAgentServer(s, grpcSrv)
	})
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create grpc transport")
	}
	httpSrv, err := NewHttpServer(r)
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create http server")
	}
	httpTransport, err := ihttp.NewServer(cfg.Http, httpSrv.Handler(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create http transport")
	}

	mgr := &Manager{
		cfg:           cfg,
		registry:      r,
		grpcSrv:       grpcSrv,
		grpcTransport: grpcTransport,
		httpSrv:       httpSrv,
		httpTransport: httpTransport,
		client:        client,
		beater:        beater,
	}
	dlog.NewStructLogger(log, mgr)
	return mgr, nil
}

// Run creates the following long running goroutines
//
// grpc server
// http server
// monitor metrics collector
// client to central server, register, keep alive
//
// short running goroutines
//
// run benchmarks
// install packages
// send metrics
//
// different go routines communicate using event bus
func (mgr *Manager) Run() error {
	var (
		wg      sync.WaitGroup
		grpcErr error
		httpErr error
		merr    = errors.NewMultiErrSafe()
	)
	wg.Add(3) // grpc + http + beater TODO: mon
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// grpc server
	go func() {
		go func() {
			if err := mgr.grpcTransport.Run(); err != nil {
				grpcErr = err
				cancel()
			}
		}()
		select {
		case <-ctx.Done():
			if grpcErr != nil {
				merr.Append(grpcErr)
				mgr.log.Errorf("can't run grpc server %v", grpcErr)
			} else {
				// other service's fault ...
				mgr.log.Warn("TODO: other's fault, need to shutdown grpc server")
			}
			wg.Done()
			return
		}
	}()
	// http server
	go func() {
		go func() {
			if err := mgr.httpTransport.Run(); err != nil {
				httpErr = err
				cancel()
			}
		}()
		select {
		case <-ctx.Done():
			if httpErr != nil {
				merr.Append(httpErr)
				mgr.log.Errorf("can't run http server %v", httpErr)
			} else {
				// other service's fault
				mgr.log.Warn("TODO: other's fault, need to shutdown http server")
			}
			wg.Done()
			return
		}
	}()
	// heartbeat with server
	go func() {
		// TODO: logic here might be incorrect, beater can exit if ctx is canceled by other go routine, i.e. grpc, http server
		if err := mgr.beater.RunWithContext(ctx); err != nil {
			merr.Append(err)
			mgr.log.Warnf("can't run beater %v", err)
			cancel()
		}
		wg.Done()
	}()

	wg.Wait()
	return merr.ErrorOrNil()
}
