package job

import (
	"context"
	"testing"
	"time"

	"github.com/dyweb/gommon/util/testutil"
	asst "github.com/stretchr/testify/assert"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
	"github.com/benchhub/benchhub/pkg/central/scheduler"
)

func TestManager_Start(t *testing.T) {
	assert := asst.New(t)

	mgr := NewManager()

	// FIXME: this is copied from TestPlanner_Job
	var job pb.JobSpec
	testutil.ReadYAMLToStrict(t, "../../../example/pingpong/pingpong.yml", &job)
	var nodes1Loader1Db []pb.Node
	testutil.ReadYAMLToStrict(t, "testdata/nodes_1l1d.yml", &nodes1Loader1Db)
	s := scheduler.NewDbBench()
	assigned, err := s.AssignNode(nodes1Loader1Db, job.NodeAssignments)
	assert.Nil(err)
	assert.Equal(2, len(assigned))

	mgr.SetSpec(job)
	mgr.SetAssignedNodes(assigned)
	err = mgr.Plan()
	assert.Nil(err)
	err = mgr.Start(context.Background())
	assert.Nil(err)
	time.Sleep(time.Millisecond) // NOTE: sleep to see the log of job manager is running
}
