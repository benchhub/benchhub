package server

import (
	"context"
	"sync"

	igrpc "github.com/at15/go.ice/ice/transport/grpc"
	ihttp "github.com/at15/go.ice/ice/transport/http"
	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"
	"google.golang.org/grpc"

	"github.com/benchhub/benchhub/pkg/central/store/meta"
	mygrpc "github.com/benchhub/benchhub/pkg/central/transport/grpc"
	"github.com/benchhub/benchhub/pkg/config"
)

type runnable interface {
	RunWithContext(ctx context.Context) error
}

type Manager struct {
	cfg config.CentralServerConfig

	registry *Registry

	meta          meta.Provider
	job           *JobPoller
	grpcSrv       *GrpcServer
	grpcTransport *igrpc.Server
	httpSrv       *HttpServer
	httpTransport *ihttp.Server

	log *dlog.Logger
}

func NewManager(cfg config.CentralServerConfig) (*Manager, error) {
	log.Infof("creating benchhub central manager")
	metaStore, err := meta.GetProvider(cfg.Meta.Provider)
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create meta store")
	}

	// registry
	r := NewRegistry(cfg)
	r.Meta = metaStore

	// job poller
	job, err := NewJobPoller(r, cfg.Job.PollInterval)
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create job controller")
	}

	// grpc http
	grpcSrv, err := NewGrpcServer(metaStore, r)
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create grpc server")
	}
	grpcTransport, err := igrpc.NewServer(cfg.Grpc, func(s *grpc.Server) {
		mygrpc.RegisterBenchHubCentralServer(s, grpcSrv)
	})
	if err != nil {
		return nil, errors.Wrap(err, "manager can't create grpc transport")
	}
	httpSrv, err := NewHttpServer(metaStore, r)
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
		meta:          metaStore,
		job:           job,
		grpcSrv:       grpcSrv,
		grpcTransport: grpcTransport,
		httpSrv:       httpSrv,
		httpTransport: httpTransport,
	}
	dlog.NewStructLogger(log, mgr)
	return mgr, nil
}

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

	names = append(names, "job-poller")
	routines = append(routines, mgr.job)

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
