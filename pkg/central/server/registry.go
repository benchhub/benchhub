package server

import (
	"sync"

	"github.com/benchhub/benchhub/pkg/central/job"
	"github.com/benchhub/benchhub/pkg/central/scheduler"
	"github.com/benchhub/benchhub/pkg/central/store/meta"
	"github.com/benchhub/benchhub/pkg/config"
)

type Registry struct {
	mu sync.RWMutex

	Config    config.CentralServerConfig
	Meta      meta.Provider
	Scheduler scheduler.Scheduler
	jobs      map[string]*job.Manager
}

func NewRegistry(cfg config.CentralServerConfig) *Registry {
	r := &Registry{
		Config:    cfg,
		Scheduler: scheduler.NewDbBench(),
		jobs:      make(map[string]*job.Manager),
	}
	return r
}

func (r *Registry) AddJob(job *job.Manager) error {
	r.mu.Lock()
	// TODO: check job id
	r.jobs[job.Id()] = job
	r.mu.Unlock()
	return nil
}
