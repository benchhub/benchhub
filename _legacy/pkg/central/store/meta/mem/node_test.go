package mem

import (
	"testing"

	asst "github.com/stretchr/testify/assert"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

func TestMetaStore_AddNode(t *testing.T) {
	assert := asst.New(t)

	m := NewMetaStore()
	m.AddNode("1", pb.Node{Info: pb.NodeInfo{Host: "n1"}})
	n, err := m.NumNodes()
	assert.Nil(err)
	assert.Equal(1, n)

	node, err := m.FindNodeById("1")
	assert.Nil(err)
	assert.Equal("n1", node.Info.Host)

	assert.NotNil(m.RemoveNode("2"))
	assert.Nil(m.RemoveNode("1"))
	n, err = m.NumNodes()
	assert.Equal(0, n)

}
