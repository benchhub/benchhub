package nodeutil

import (
	"testing"

	"github.com/benchhub/benchhub/pkg/common/config"
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

func TestGetNodeInfo(t *testing.T) {
	assert := asst.New(t)
	n, err := GetNodeInfo(config.NodeConfig{})
	assert.Nil(err)
	t.Logf("start time %d", n.StartTime)
	t.Logf("boot  time %d", n.BootTime)
	t.Logf("cores %d", n.Capacity.Cores)
	t.Logf("disk total %d MB", n.Capacity.DiskTotal)
	t.Logf("disk free %d MB", n.Capacity.DiskFree)
	t.Logf("mem total %d MB", n.Capacity.MemoryTotal)
	t.Logf("mem free %d MB", n.Capacity.MemoryFree)
}
