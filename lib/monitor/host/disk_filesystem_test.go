package host

import (
	"testing"

	asst "github.com/stretchr/testify/assert"
)

func TestFilesystem_Update(t *testing.T) {
	assert := asst.New(t)

	fs := NewFilesystem("")
	assert.Nil(fs.Update())
	//"github.com/davecgh/go-spew/spew"
	//spew.Dump(*fs)
	// print in gb
	t.Logf("block size %d", fs.BlockSize)
	t.Logf("total %dGB", fs.Blocks*4/1024/1024)
	t.Logf("free %dGB", fs.BlocksFree*4/1024/1024)
	t.Logf("avail %dGB", fs.BlocksAvail*4/1024/1024)
}
