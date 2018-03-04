package server

import "github.com/benchhub/benchhub/pkg/agent/config"

// Registry is a central repository for shared states, i.e. data store, metrics etc.
type Registry struct {
	Config config.ServerConfig
	State  *StateMachine
}
