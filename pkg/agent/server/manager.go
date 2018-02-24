package server

import (
	"context"
	"sync"

	dlog "github.com/dyweb/gommon/log"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	igrpc "github.com/at15/go.ice/ice/transport/grpc"
	ihttp "github.com/at15/go.ice/ice/transport/http"
	"github.com/benchhub/benchhub/pkg/agent/config"
	mygrpc "github.com/benchhub/benchhub/pkg/agent/transport/grpc"
	myhttp "github.com/benchhub/benchhub/pkg/agent/transport/http"
)

type Manager struct {
	cfg           config.ServerConfig
	registry      *Registry
	grpcSrv       *GrpcServer
	grpcTransport *igrpc.Server
	httpSrv       *HttpServer
	httpTransport *ihttp.Server
	log           *dlog.Logger
}

func NewManager(cfg config.ServerConfig) (*Manager, error) {
	log.Info("creating benchhub agent manager")
	grpcSrv, err := NewGrpcServer()
	if err != nil {
		return nil, errors.WithMessage(err, "manager can't create grpc server")
	}
	grpcTransport, err := igrpc.NewServer(cfg.Grpc, func(s *grpc.Server) {
		mygrpc.RegisterBenchHubAgentServer(s, grpcSrv)
	})
	if err != nil {
		return nil, errors.WithMessage(err, "manager can't create grpc transport")
	}
	httpSrv, err := NewHttpServer()
	if err != nil {
		return nil, errors.WithMessage(err, "manager can't create http server")
	}
	httpHandler := myhttp.NewHttpHandler(httpSrv.HandlerRegister)
	httpTransport, err := ihttp.NewServer(cfg.Http, httpHandler, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "manager can't create http transport")
	}
	mgr := &Manager{
		cfg:           cfg,
		grpcSrv:       grpcSrv,
		grpcTransport: grpcTransport,
		httpSrv:       httpSrv,
		httpTransport: httpTransport,
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
		grpcErr error // TODO: multiple error
		httpErr error
	)
	wg.Add(2)
	ctx, cancel := context.WithCancel(context.Background())
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
				mgr.log.Errorf("can't run http server %v", httpErr)
			} else {
				mgr.log.Warn("TODO: other's fault, need to shutdown http server")
			}
			wg.Done()
			return
		}
	}()
	wg.Wait()
	return grpcErr
}
