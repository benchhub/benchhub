package server

import (
	"sync"

	"github.com/benchhub/benchhub/pkg/central/config"
	"github.com/benchhub/benchhub/pkg/central/job"
	"github.com/benchhub/benchhub/pkg/central/store/meta"
)

type Registry struct {
	mu sync.RWMutex

	Config config.ServerConfig
	Meta   meta.Provider
	jobs   map[string]*job.Manager
}

func NewRegistry(cfg config.ServerConfig) *Registry {
	r := &Registry{
		Config: cfg,
		jobs:   make(map[string]*job.Manager),
	}
	return r
}
