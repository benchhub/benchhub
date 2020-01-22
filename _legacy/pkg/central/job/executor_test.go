package job

import (
	"bytes"
	"context"
	"io"
	"sync"
	"testing"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
	"github.com/dyweb/gommon/util/testutil"
	asst "github.com/stretchr/testify/assert"
)

func TestEchoExecutor(t *testing.T) {
	assert := asst.New(t)

	plan := pingpongPlan(t)

	// FIXME: this test is just for early stage design
	w := &bytes.Buffer{}
	for i := range plan.Pipelines {
		excPipeline(t, plan.Pipelines[i], w)
	}
	// TODO: wait for change in gommon/testutil https://github.com/dyweb/gommon/issues/64
	golden := "testdata/echo_pingpong_result.txt"
	if testutil.GenGolden().B() {
		testutil.WriteFixture(t, golden, w.Bytes())
	} else {
		b := testutil.ReadFixture(t, golden)
		assert.Equal(string(b), w.String())
	}
}

func excPipeline(t *testing.T, plan pb.StagePipelinePlan, w io.Writer) {
	t.Logf("execute pipeline %s with %d stages", plan.Name, len(plan.Stages))
	// execute stages in one pipeline concurrently
	// TODO: this is not async ...
	var wg sync.WaitGroup
	wg.Add(len(plan.Stages))
	for i := range plan.Stages {
		go func(i int) {
			excStage(t, plan.Stages[i], w)
			t.Logf("stage %d finished", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func excStage(t *testing.T, plan pb.StagePlan, w io.Writer) {
	// create executors based on nodes
	var executors []Executor
	for i := 0; i < len(plan.Nodes); i++ {
		executors = append(executors, NewEchoExecutor(plan, i, w))
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
