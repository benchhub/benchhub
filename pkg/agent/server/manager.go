package server

import (
	"context"
	"sync"

	igrpc "github.com/at15/go.ice/ice/transport/grpc"
	ihttp "github.com/at15/go.ice/ice/transport/http"
	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"
	"google.golang.org/grpc"

	mygrpc "github.com/benchhub/benchhub/pkg/agent/transport/grpc"
	crpc "github.com/benchhub/benchhub/pkg/central/transport/grpc"
	"github.com/benchhub/benchhub/pkg/config"
)

type runnable interface {
	RunWithContext(ctx context.Context) error
}

type Manager struct {
	cfg config.AgentServerConfig

	registry *Registry

	grpcSrv       *GrpcServer
	grpcTransport *igrpc.Server
	httpSrv       *HttpServer
	httpTransport *ihttp.Server

	client crpc.BenchHubCentralClient
	beater *Beater
	log    *dlog.Logger
}

func NewManager(cfg config.AgentServerConfig) (*Manager, error) {
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

	// grpc
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
	// http
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
// beater, register node, send heartbeat
// grpc server
// http server
// TODO: metrics collector
func (mgr *Manager) Run() error {
	var (
		wg   sync.WaitGroup
		merr = errors.NewMultiErrSafe()
	)
	names, routines := mgr.runnable()
	wg.Add(len(routines))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for i, name := range names {
		go func(i int, name string) {
			log.Infof("%s started", name)
			err := routines[i].RunWithContext(ctx)
			log.Infof("%s stopped", name)
			wg.Done()
			if err != nil {
				log.Errorf("%s error %v", name, err)
				log.Errorf("cancel manager context due to %s", name)
				cancel()
			}
		}(i, name)
	}
	wg.Wait()
	return merr.ErrorOrNil()
}

// NOTE: we return two array instead of a map because iterate map has random order
func (mgr *Manager) runnable() ([]string, []runnable) {
	var names []string
	var routines []runnable

	names = append(names, "beater")
	routines = append(routines, mgr.beater)

	names = append(names, "grpc-server")
	routines = append(routines, mgr.grpcTransport)
	names = append(names, "http-server")
	routines = append(routines, mgr.httpTransport)

	if len(names) != len(routines) {
		// TODO: need an assert package
		panic("length of names and routines does not equal")
	}

	return names, routines
}
