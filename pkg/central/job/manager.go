package job

import (
	"context"
	"sync"

	dlog "github.com/dyweb/gommon/log"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

type Manager struct {
	mu sync.RWMutex

	spec      pb.JobSpec
	nodes     []pb.AssignedNode
	planner   *Planner
	executors []Executor
	log       *dlog.Logger
}

func NewManager() *Manager {
	c := &Manager{
		planner: NewPlanner(),
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

// Start creates a goroutine and return immediately
func (m *Manager) Start() error {

	return nil
}

// Run block and return until it is finished
func (m *Manager) Run(ctx context.Context) error {
	return nil
}
