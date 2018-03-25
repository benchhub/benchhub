package job

import (
	"context"
	"sync"

	dlog "github.com/dyweb/gommon/log"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

type ExecutorStatus uint8

const (
	ExecutorUnknown = iota
	ExecutorRunning
	ExecutorFinished
	ExecutorError
)

// Executor runs plan on a single node, coordination between nodes is done by manager
// - simply print command out (for dry run)
// - run everything locally, assume the job spec can be run locally, i.e. everything is localhost
// - dispatch job to remote agent
type Executor interface {
	rwMutex
	// Start dispatch/run the plan but does not wait for it to complete
	Start(ctx context.Context) error
	Status() (ExecutorStatus, error)
}

type rwMutex interface {
	RLock()
	RUnlock()
	Lock()
	Unlock()
}

//var _ Executor = (*MockExecutor)(nil)
var _ Executor = (*EchoExecutor)(nil)

//var _ Executor = (*LocalExecutor)(nil)

type MockExecutor struct {
	*sync.RWMutex

	log *dlog.Logger
}

func (exc *MockExecutor) Start(ctx context.Context) error {
	return nil
}

type EchoExecutor struct {
	*sync.RWMutex

	plan pb.StagePlan
	log  *dlog.Logger
}

func NewEchoExecutor(plan pb.StagePlan) *EchoExecutor {
	exc := &EchoExecutor{
		plan: plan,
	}
	dlog.NewStructLogger(log, exc)
	return exc
}

func (exc *EchoExecutor) Start(ctx context.Context) error {
	for _, p := range exc.plan.Pipelines {
		for _, t := range p.Tasks {
			exc.log.Infof("task driver %s", t.Spec.Driver)
			switch t.Spec.Driver {
			case pb.TaskDriver_SHELL:
				exc.log.Infof("shell %s", t.Spec.Shell.Command)
			case pb.TaskDriver_EXEC:
				exc.log.Infof("exec %s %s", t.Spec.Exec.Command, t.Spec.Exec.Args)
			case pb.TaskDriver_DOCKER:
				exc.log.Infof("docker %s %s", t.Spec.Docker.Image, t.Spec.Docker.Action)
			}
		}
	}
	return nil
}

func (exc *EchoExecutor) Status() (ExecutorStatus, error) {
	return ExecutorFinished, nil
}

type LocalExecutor struct {
	*sync.RWMutex

	log *dlog.Logger
}

func (exc *LocalExecutor) Start(ctx context.Context) error {
	return nil
}
