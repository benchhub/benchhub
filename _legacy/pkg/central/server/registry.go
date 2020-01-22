package server

import (
	"sync"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
	"github.com/benchhub/benchhub/pkg/central/job"
	"github.com/benchhub/benchhub/pkg/central/scheduler"
	"github.com/benchhub/benchhub/pkg/central/store/meta"
	"github.com/benchhub/benchhub/pkg/config"
	"github.com/benchhub/benchhub/pkg/util/nodeutil"
)

type Registry struct {
	mu sync.RWMutex

	Config    config.CentralServerConfig
	Meta      meta.Provider
	Scheduler scheduler.Scheduler
	jobs      map[string]*job.Manager
	nodeInfo  pb.NodeInfo
}

func NewRegistry(cfg config.CentralServerConfig) (*Registry, error) {
	r := &Registry{
		Config:    cfg,
		Scheduler: scheduler.NewDbBench(),
		jobs:      make(map[string]*job.Manager),
	}
	info, err := nodeutil.GetNodeInfo(cfg.Node, cfg.Grpc.Addr)
	if err != nil {
		return nil, err
	}
	r.nodeInfo = *info
	return r, nil
}

func (r *Registry) AddJob(job *job.Manager) error {
	r.mu.Lock()
	// TODO: check job id
	r.jobs[job.Id()] = job
	r.mu.Unlock()
	return nil
}

func (r *Registry) NodeInfo() pb.NodeInfo {
	return r.nodeInfo
}
