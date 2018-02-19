package host

import (
	"testing"

	asst "github.com/stretchr/testify/assert"
)

func TestBlockDevices_Update(t *testing.T) {
	assert := asst.New(t)

	devices := NewBlockDevices("testdata/proc/diskstats")
	assert.Nil(devices.Update())

	assert.Equal("loop0", devices.Devices["loop0"].Name)
	assert.Equal(uint64(148503), devices.Devices["sda"].ReadsCompleted)
	assert.Equal(uint64(190248), devices.Devices["sda"].TimeSpentWritingMs)

	devices = NewBlockDevices("")
	assert.Nil(devices.Update())
	// TODO: maybe we should use slice of struct instead of slice of pointer to struct, latter makes default print easier
	t.Log(devices.Devices)
}
