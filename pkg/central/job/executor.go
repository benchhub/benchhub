package job

import (
	"context"
	"fmt"
	"io"
	"sync"

	dlog "github.com/dyweb/gommon/log"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

type ExecutorStatus uint8

const (
	ExecutorUnknown = iota
	ExecutorIdle
	ExecutorRunning
	ExecutorFinished
	ExecutorError
)

// Executor runs a stage plan on a single node, coordination between nodes is done by manager
// - simply print command out (for dry run)
// - run everything locally, assume the job spec can be run locally, i.e. everything is localhost
// - dispatch job to remote agent
type Executor interface {
	// Start dispatch/run the plan but does not wait for it to complete
	Start(ctx context.Context) error
	Status() (ExecutorStatus, error)
}

//var _ Executor = (*MockExecutor)(nil)
var _ Executor = (*EchoExecutor)(nil)

//var _ Executor = (*LocalExecutor)(nil)

type MockExecutor struct {
	mu sync.RWMutex

	log *dlog.Logger
}

func (exc *MockExecutor) Start(ctx context.Context) error {
	return nil
}

// EchoExecutor output command it should run to io.Writer
type EchoExecutor struct {
	mu sync.RWMutex

	node   pb.AssignedNode
	plan   pb.StagePlan
	status ExecutorStatus
	w      io.Writer

	log *dlog.Logger
}

func NewEchoExecutor(plan pb.StagePlan, nodeIndex int, w io.Writer) *EchoExecutor {
	exc := &EchoExecutor{
		node:   plan.Nodes[nodeIndex],
		plan:   plan,
		status: ExecutorIdle,
		w:      w,
	}
	dlog.NewStructLogger(log, exc)
	return exc
}

func (exc *EchoExecutor) Start(ctx context.Context) error {
	exc.mu.Lock()
	exc.status = ExecutorRunning
	exc.mu.Unlock()
	fmt.Fprintf(exc.w, "execute stage %s on node %s\n", exc.plan.Name, exc.node.Spec.Name)
	for _, p := range exc.plan.Pipelines {
		for _, t := range p.Tasks {
			fmt.Fprintf(exc.w, "task driver %s\n", t.Spec.Driver)
			switch t.Spec.Driver {
			case pb.TaskDriver_SHELL:
				fmt.Fprintf(exc.w, "shell %s\n", t.Spec.Shell.Command)
			case pb.TaskDriver_EXEC:
				fmt.Fprintf(exc.w, "exec %s %s\n", t.Spec.Exec.Command, t.Spec.Exec.Args)
			case pb.TaskDriver_DOCKER:
				fmt.Fprintf(exc.w, "docker %s %s\n", t.Spec.Docker.Image, t.Spec.Docker.Action)
			}
		}
	}
	exc.mu.Lock()
	exc.status = ExecutorFinished
	exc.mu.Unlock()
	return nil
}

func (exc *EchoExecutor) Status() (ExecutorStatus, error) {
	exc.mu.RLock()
	s := exc.status
	exc.mu.RUnlock()
	return s, nil
}

type LocalExecutor struct {
	*sync.RWMutex

	log *dlog.Logger
}

func (exc *LocalExecutor) Start(ctx context.Context) error {
	return nil
}
