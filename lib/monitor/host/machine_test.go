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
}
