package mem

import (
	"testing"

	asst "github.com/stretchr/testify/assert"

	pbc "github.com/benchhub/benchhub/pkg/common/commonpb"
)

func TestMetaStore_AddNode(t *testing.T) {
	assert := asst.New(t)

	m := NewMetaStore()
	m.AddNode("1", pbc.Node{Host: "n1"})
	n, err := m.NumNodes()
	assert.Nil(err)
	assert.Equal(1, n)

	node, err := m.FindNodeById("1")
	assert.Nil(err)
	assert.Equal("n1", node.Host)

	assert.NotNil(m.RemoveNode("2"))
	assert.Nil(m.RemoveNode("1"))
	n, err = m.NumNodes()
	assert.Equal(0, n)

}
