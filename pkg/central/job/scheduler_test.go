package job

import (
	"testing"

	asst "github.com/stretchr/testify/assert"
	pb "github.com/benchhub/benchhub/pkg/bhpb"
	"github.com/dyweb/gommon/util/testutil"
)

func TestScheduler_AssignNode(t *testing.T) {
	assert := asst.New(t)
	s := NewScheduler()
	var spec1Loader1Db []pb.NodeAssignmentSpec
	testutil.ReadYAMLToStrict(t, "testdata/nodesassign_1l1d.yml", &spec1Loader1Db)
	assert.Equal(2, len(spec1Loader1Db))
	//t.Log(spec1Loader1Db)
	t.Run("no node", func(t *testing.T) {
		assert := asst.New(t)
		_, err := s.AssignNode(nil, spec1Loader1Db)
		assert.NotNil(err)
		assert.Equal("0 nodes available", err.Error())
	})

	// exact match
	t.Run("1l1d", func(t *testing.T) {
		assert := asst.New(t)

		var nodes1Loader1Db []pb.Node
		testutil.ReadYAMLToStrict(t, "testdata/nodes_1l1d.yml", &nodes1Loader1Db)
		t.Log(len(nodes1Loader1Db))
		t.Log(nodes1Loader1Db[0].Info.Role)
		res, err := s.AssignNode(nodes1Loader1Db, spec1Loader1Db)
		assert.Nil(err)
		assert.Equal(2, len(res))
		// FIXME: we can assert this because AssignNode is using map which iter is random
		//assert.Equal("srv", res[0].Spec.Name)
		//assert.Equal("cli", res[1].Spec.Name)
	})
}
