package server

import (
	"context"
	"sync"

	igrpc "github.com/at15/go.ice/ice/transport/grpc"
	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"
	"google.golang.org/grpc"

	"github.com/benchhub/benchhub/pkg/central/config"
	"github.com/benchhub/benchhub/pkg/central/store/meta"
	mygrpc "github.com/benchhub/benchhub/pkg/central/transport/grpc"
)

type Manager struct {
	cfg config.ServerConfig

	meta          meta.Provider
	grpcSrv       *GrpcServer
	grpcTransport *igrpc.Server

	log *dlog.Logger
}

func NewManager(cfg config.ServerConfig) (*Manager, error) {
	log.Infof("creating benchhub central manager")
	metaStore, err := meta.GetProvider(cfg.Meta.Provider)
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create meta store")
	}
	grpcSrv, err := NewGrpcServer()
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create grpc server")
	}
	grpcTransport, err := igrpc.NewServer(cfg.Grpc, func(s *grpc.Server) {
		mygrpc.RegisterBenchHubCentralServer(s, grpcSrv)
	})
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create grpc transport")
	}
	mgr := &Manager{
		cfg:           cfg,
		meta:          metaStore,
		grpcSrv:       grpcSrv,
		grpcTransport: grpcTransport,
	}
	dlog.NewStructLogger(log, mgr)
	return mgr, nil
}

func (mgr *Manager) Run() error {
	var (
		wg      sync.WaitGroup
		grpcErr error
		merr    = errors.NewMultiErrSafe()
	)
	wg.Add(1) // grpc + TODO: http
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
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
				mgr.log.Warn("TODO: other's fault, need to shutdown grpc server")
			}
			wg.Done()
			return
		}
	}()
	wg.Wait()
	return merr.ErrorOrNil()
}
