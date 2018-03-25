package job

import (
	"context"
	"testing"
)

func TestExecutor(t *testing.T) {
	plan := pingpongPlan(t)

	// FIXME: this test if just for early stage design
	pipeline := plan.Pipelines[0]
	t.Logf("execute pipeline %s", pipeline.Name)
	// for each stage, assign # executors that equals to assigned nodes
	// FIXME: should execute stages in one pipeline concurrently
	stage := pipeline.Stages[0]
	t.Logf("execute stage %s", stage.Name)

	// create executors based on nodes
	var executors []Executor
	for i := 0; i < len(stage.Nodes); i++ {
		// TODO: stage plan contains all nodes, but actually the executor just need to know one node, itself
		executors = append(executors, NewEchoExecutor(stage))
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
