package host

import (
	"testing"

	asst "github.com/stretchr/testify/assert"
	"runtime"
)

func TestCpus_Update(t *testing.T) {
	assert := asst.New(t)

	cpus := NewCpus("testdata/proc/stat")
	assert.Nil(cpus.Update())
	assert.Equal(8, len(cpus.Cores))
	// NOTE: we didn't divide it by CPU USER_HZ (which is 100 in most cases)
	assert.Equal(uint64(849251), cpus.Total.User)

	cpus = NewCpus("")
	assert.Nil(cpus.Update())
	assert.Equal(runtime.NumCPU(), len(cpus.Cores))
}
