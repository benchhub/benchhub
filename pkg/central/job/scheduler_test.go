package job

import (
	"testing"

	"github.com/dyweb/gommon/util/testutil"
	asst "github.com/stretchr/testify/assert"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
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
		assert.Equal(2, len(nodes1Loader1Db))
		assert.Equal(pb.Role_LOADER, nodes1Loader1Db[0].Info.Role)
		res, err := s.AssignNode(nodes1Loader1Db, spec1Loader1Db)
		assert.Nil(err)
		assert.Equal(2, len(res))
		// the result is in the order of specification
		assert.Equal(spec1Loader1Db[0].Name, res[0].Spec.Name)
		assert.Equal(spec1Loader1Db[1].Name, res[1].Spec.Name)
		// node role is updated
		for _, r := range res {
			assert.Equal(r.Spec.Role, r.Node.Role)
		}
	})
}
