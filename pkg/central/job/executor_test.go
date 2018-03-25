package job

import (
	"context"
	"sync"
	"testing"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

func TestExecutor(t *testing.T) {
	plan := pingpongPlan(t)

	// FIXME: this test is just for early stage design
	for i := range plan.Pipelines {
		excPipeline(t, plan.Pipelines[i])
	}
}

func excPipeline(t *testing.T, plan pb.StagePipelinePlan) {
	t.Logf("execute pipeline %s with %d stages", plan.Name, len(plan.Stages))
	// execute stages in one pipeline concurrently
	var wg sync.WaitGroup
	wg.Add(len(plan.Stages))
	for i := range plan.Stages {
		go func(i int) {
			excStage(t, plan.Stages[i])
			t.Logf("stage %d finished", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func excStage(t *testing.T, plan pb.StagePlan) {
	// create executors based on nodes
	var executors []Executor
	for i := 0; i < len(plan.Nodes); i++ {
		// TODO: stage plan contains all nodes, but actually the executor just need to know one node, itself
		executors = append(executors, NewEchoExecutor(plan))
	}

	// run all the executors
	for i := 0; i < len(executors); i++ {
		err := executors[i].Start(context.Background())
		if err != nil {
			t.Fatalf("failed to start executor %d %v", i, err)
		}
	}

	// check states of all executors
	for i := 0; i < len(executors); i++ {
		s, err := executors[i].Status()
		if err != nil {
			t.Fatalf("failed to get executor status %d %v", i, err)
		}
		if s == ExecutorFinished {
			t.Logf("executor %d finished", i)
		}
	}
}
