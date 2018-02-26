package host

import (
	"testing"

	asst "github.com/stretchr/testify/assert"
)

func TestMachineStat_Update(t *testing.T) {
	assert := asst.New(t)
	m := &Machine{}
	assert.Nil(m.Update())
	t.Logf("host %s", m.HostName)
	t.Logf("num cores %d", m.NumCores)
	t.Logf("disk space total %d bytes", m.DiskSpaceTotal)
	t.Logf("disk space free %d bytes", m.DiskSpaceFree)
	t.Logf("disk inode total %d", m.DiskInodeTotal)
	t.Logf("disk inode freee %d", m.DiskInodeFree)
	t.Logf("boot time %d", m.BootTime)
}
