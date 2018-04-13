package scheduler

import (
	"testing"

	tu "github.com/dyweb/gommon/util/testutil"
	asst "github.com/stretchr/testify/assert"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
	btu "github.com/benchhub/benchhub/pkg/util/testutil"
)

func TestDbBench_AssignNode(t *testing.T) {
	assert := asst.New(t)

	var spec1Loader1Db []pb.NodeAssignmentSpec
	tu.ReadYAMLToStrict(t, btu.CentralTestdata("nodesassign_1l1d.yml"), &spec1Loader1Db)
	assert.Equal(2, len(spec1Loader1Db))
	//t.Log(spec1Loader1Db)

	s := NewDbBench()
	t.Run("no node", func(t *testing.T) {
		assert := asst.New(t)
		_, err := s.AssignNode(nil, spec1Loader1Db)
		assert.NotNil(err)
		assert.Equal("0 nodes available", err.Error())
	})

	// exact match
	//t.Run("1l1d", func(t *testing.T) {
	//	assert := asst.New(t)
	//
	//	var nodes1Loader1Db []pb.Node
	//	tu.ReadYAMLToStrict(t, btu.CentralTestdata("nodes_1l1d.yml"), &nodes1Loader1Db)
	//	assert.Equal(2, len(nodes1Loader1Db))
	//	//assert.Equal(pb.Role_LOADER, nodes1Loader1Db[0].Info.Role)
	//
	//	res, err := s.AssignNode(nodes1Loader1Db, spec1Loader1Db)
	//	assert.Nil(err)
	//
	//	assert.Equal(2, len(res))
	//	// the result is in the order of specification
	//	assert.Equal(spec1Loader1Db[0].Properties.Name, res[0].Spec.Properties.Name)
	//	assert.Equal(spec1Loader1Db[1].Properties.Name, res[1].Spec.Properties.Name)
	//})
	//
	//t.Run("2a", func(t *testing.T) {
	//	assert := asst.New(t)
	//
	//	var nodes2a []pb.Node
	//	tu.ReadYAMLToStrict(t, btu.CentralTestdata("nodes_2a.yml"), &nodes2a)
	//	assert.Equal(2, len(nodes2a))
	//	//assert.Equal(pb.Role_ANY, nodes2a[0].Info.Role)
	//	//assert.Equal(pb.Role_ANY, nodes2a[1].Info.Role)
	//
	//	res, err := s.AssignNode(nodes2a, spec1Loader1Db)
	//	assert.Nil(err)
	//
	//	assert.Equal(2, len(res))
	//})

}
