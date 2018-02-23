package host

import (
	"testing"

	asst "github.com/stretchr/testify/assert"
)

func TestBlockDevices_Update(t *testing.T) {
	assert := asst.New(t)

	devices := NewBlockDevices("testdata/proc/diskstats")
	assert.Nil(devices.Update())

	assert.Equal("loop0", devices.Devices[0].Name)
	assert.Equal("sda", devices.Devices[8].Name)
	assert.Equal(uint64(148503), devices.Devices[8].ReadsCompleted)
	assert.Equal(uint64(190248), devices.Devices[8].TimeSpentWritingMs)

	devices = NewBlockDevices("")
	assert.Nil(devices.Update())
	numDevices := len(devices.Devices)
	assert.Nil(devices.Update())
	assert.Equal(numDevices, len(devices.Devices))
	//"github.com/davecgh/go-spew/spew"
	//spew.Dump(devices.Devices)
}
