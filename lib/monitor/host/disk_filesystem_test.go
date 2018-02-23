package host

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	asst "github.com/stretchr/testify/assert"
)

func TestFilesystem_Update(t *testing.T) {
	assert := asst.New(t)

	fs := NewFilesystem("")
	assert.Nil(fs.Update())
	spew.Dump(*fs)
	// print in gb
	t.Logf("total %dGB", fs.Blocks*4/1024/1024)
	t.Logf("free %dGB", fs.BlocksFree*4/1024/1024)
	t.Logf("avail %dGB", fs.BlocksAvail*4/1024/1024)
}
