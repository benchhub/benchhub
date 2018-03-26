package job

import (
	"context"
	"sync"

	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

type Manager struct {
	mu sync.RWMutex

	spec      *pb.JobSpec
	plan      *pb.JobPlan
	nodes     []pb.AssignedNode
	executors []Executor
	planner   *Planner

	ctx    context.Context
	cancel context.CancelFunc

	log *dlog.Logger
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
	m.spec = &spec
	m.mu.Unlock()
}

func (m *Manager) SetAssignedNodes(nodes []pb.AssignedNode) {
	m.mu.Lock()
	m.nodes = nodes
	m.mu.Unlock()
}

// Plan generate plan based on job spec
func (m *Manager) Plan() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	plan, err := m.planner.Job(m.nodes, *m.spec)
	if err != nil {
		return errors.Wrap(err, "manager failed to plan")
	}
	m.plan = &plan
	return nil
}

// Start creates a goroutine and return immediately, error only comes from incorrect configuration
// MUST call Plan before call Start
func (m *Manager) Start(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// check spec
	if m.spec == nil || len(m.spec.Pipelines) == 0 || len(m.spec.Stages) == 0 {
		return errors.New("empty pipelines or stages, is SetSpec called?")
	}
	// check nodes
	if len(m.nodes) == 0 {
		return errors.New("empty nodes, is SetAssignedNodes called?")
	}
	// check plan
	if m.plan == nil || len(m.plan.Pipelines) == 0 {
		return errors.New("empty plan, is Plan called?")
	}
	m.ctx, m.cancel = context.WithCancel(ctx)
	go func() {
		waitCh := make(chan error)
		go func() {
			if err := m.run(); err != nil {
				waitCh <- err
			}
		}()
		select {
		case err := <-waitCh:
			m.log.Errorf("manager run failed %v", err)
			return
		case <-m.ctx.Done():
			m.log.Info("manager canceled by context")
			return
		}
		// TODO: set manager status here
	}()
	m.log.Info("job manager started")
	return nil
}

func (m *Manager) run() error {
	// TODO: maybe just take the lock and release until the end? no ... what if the manger is cancelled
	// TODO: start cron to dispatch etc.
	m.log.Info("job manager is running")
	return nil
}
