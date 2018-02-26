package nodeutil

import (
	"testing"

	"github.com/rs/xid"
	asst "github.com/stretchr/testify/assert"
)

func TestNewUID(t *testing.T) {
	assert := asst.New(t)
	s := NewUID()
	id, err := xid.FromString(s)
	assert.Nil(err)
	assert.False(id.IsNil())
	t.Log(id.Time())
	// TODO: print pid, it's uint64 in big endian
}
