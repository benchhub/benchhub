package job

import (
	"os"
	"testing"

	"github.com/dyweb/gommon/util/testutil"
	asst "github.com/stretchr/testify/assert"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
	"github.com/benchhub/benchhub/pkg/central/scheduler"
)

func TestPlanner_Job(t *testing.T) {
	assert := asst.New(t)

	var job pb.JobSpec
	testutil.ReadYAMLToStrict(t, "../../../example/pingpong/pingpong.yml", &job)

	var nodes1Loader1Db []pb.Node
	testutil.ReadYAMLToStrict(t, "testdata/nodes_1l1d.yml", &nodes1Loader1Db)

	s := scheduler.NewDbBench()
	assigned, err := s.AssignNode(nodes1Loader1Db, job.NodeAssignments)
	assert.Nil(err)
	assert.Equal(2, len(assigned))

	p := NewPlanner()
	plan, err := p.Job(assigned, job)
	assert.Nil(err)

	if testutil.Dump().B() {
		testutil.PrintTidyJsonTo(t, plan, os.Stderr)
	}
}

// TODO: should be able to disable planner logging, it is not logging at info level
func pingpongPlan(t *testing.T) pb.JobPlan {
	var job pb.JobSpec
	testutil.ReadYAMLToStrict(t, "../../../example/pingpong/pingpong.yml", &job)

	var nodes1Loader1Db []pb.Node
	testutil.ReadYAMLToStrict(t, "testdata/nodes_1l1d.yml", &nodes1Loader1Db)

	s := scheduler.NewDbBench()
	assigned, err := s.AssignNode(nodes1Loader1Db, job.NodeAssignments)
	if err != nil {
		t.Fatal(err)
	}
	if len(assigned) != 2 {
		t.Fatalf("expect assigned node lenght to be 2 but got %d", len(assigned))
	}

	p := NewPlanner()
	plan, err := p.Job(assigned, job)
	if err != nil {
		t.Fatal(err)
	}

	if testutil.Dump().B() {
		testutil.PrintTidyJsonTo(t, plan, os.Stderr)
	}
	return plan
}
