package job

import (
	"sync"

	dlog "github.com/dyweb/gommon/log"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

type Manager struct {
	mu sync.RWMutex

	spec     pb.JobSpec
	nodes    []pb.AssignedNode
	planner  *Planner
	executor *Executor
	log      *dlog.Logger
}

func NewManager() *Manager {
	c := &Manager{
		planner:  NewPlanner(),
		executor: NewExecutor(),
	}
	dlog.NewStructLogger(log, c)
	return c
}

func (m *Manager) Id() string {
	return m.spec.Id
}

func (m *Manager) SetSpec(spec pb.JobSpec) {
	m.mu.Lock()
	m.spec = spec
	m.mu.Unlock()
}

func (m *Manager) SetAssignedNodes(nodes []pb.AssignedNode) {
	m.mu.Lock()
	m.nodes = nodes
	m.mu.Unlock()
}
