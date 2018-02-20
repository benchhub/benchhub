package host

import (
	"testing"

	asst "github.com/stretchr/testify/assert"
)

func TestMem_Update(t *testing.T) {
	assert := asst.New(t)
	mem := NewMem("testdata/proc/meminfo")
	assert.Nil(mem.Update())
	assert.Equal(uint64(32853904), mem.MemTotal)
	assert.Equal(uint64(117620), mem.PageTables)

	mem = NewMem("")
	assert.Nil(mem.Update())
	assert.True(mem.MemTotal > 0)
	assert.True(mem.PageTables > 0)
	// TODO: why total != free + mem ... we are parsing the file correctly, it's just I don't understand the numbers
	//assert.Equal(uint64(mem.MemTotal), mem.MemFree+mem.MemAvailable)
}
