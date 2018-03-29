package server

import (
	"sync"

	"github.com/benchhub/benchhub/pkg/agent/job"
	"github.com/benchhub/benchhub/pkg/config"
)

// Registry is a central repository for shared states, i.e. data store, metrics etc.
type Registry struct {
	mu sync.RWMutex

	Config config.AgentServerConfig
	State  *StateMachine
	jobs   map[string]*job.Manager
}

func NewRegistry(cfg config.AgentServerConfig) *Registry {
	r := &Registry{
		Config: cfg,
		jobs:   make(map[string]*job.Manager),
	}
	return r
}
