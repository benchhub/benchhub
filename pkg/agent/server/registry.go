package server

import (
	"sync"

	"github.com/benchhub/benchhub/pkg/agent/job"
	pb "github.com/benchhub/benchhub/pkg/bhpb"
	"github.com/benchhub/benchhub/pkg/config"
	"github.com/benchhub/benchhub/pkg/util/nodeutil"
)

// Registry is a central repository for shared states, i.e. data store, metrics etc.
type Registry struct {
	mu sync.RWMutex

	Config   config.AgentServerConfig
	State    *StateMachine
	jobs     map[string]*job.Manager
	nodeInfo pb.NodeInfo
}

func NewRegistry(cfg config.AgentServerConfig) (*Registry, error) {
	r := &Registry{
		Config: cfg,
		jobs:   make(map[string]*job.Manager),
	}
	info, err := nodeutil.GetNodeInfo(cfg.Node, cfg.Grpc.Addr)
	if err != nil {
		return nil, err
	}
	r.nodeInfo = *info
	return r, nil
}

func (r *Registry) NodeInfo() pb.NodeInfo {
	return r.nodeInfo
}
