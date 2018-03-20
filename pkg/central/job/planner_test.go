package job

import (
	"testing"

	"github.com/dyweb/gommon/util/testutil"
	asst "github.com/stretchr/testify/assert"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
	"os"
)

func TestPlanner_Job(t *testing.T) {
	assert := asst.New(t)

	var job pb.JobSpec
	testutil.ReadYAMLToStrict(t, "../../../example/pingpong/pingpong.yml", &job)

	var nodes1Loader1Db []pb.Node
	testutil.ReadYAMLToStrict(t, "testdata/nodes_1l1d.yml", &nodes1Loader1Db)

	s := NewScheduler()
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
