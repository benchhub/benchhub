package nodeutil

import (
	"testing"

	"github.com/rs/xid"
	asst "github.com/stretchr/testify/assert"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
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
	n, err := GetNodeInfo(pb.NodeConfig{}, ":6081")
	assert.Nil(err)
	t.Logf("start time %d", n.Property.StartTime)
	t.Logf("boot  time %d", n.Property.BootTime)
	t.Logf("cores %d", n.Capacity.Cores)
	t.Logf("disk total %d MB", n.Capacity.DiskTotal)
	t.Logf("disk free  %d MB", n.Capacity.DiskFree)
	t.Logf("mem total  %d MB", n.Capacity.MemoryTotal)
	t.Logf("mem free   %d MB", n.Capacity.MemoryFree)
}
