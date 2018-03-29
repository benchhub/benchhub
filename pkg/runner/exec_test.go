package runner

import (
	"context"
	"testing"
	"time"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
	asst "github.com/stretchr/testify/assert"
)

func TestExec_Run(t *testing.T) {
	if testing.Short() {
		t.Skipf("skip runner test in short tests")
	}
	t.Run("echo", func(t *testing.T) {
		assert := asst.New(t)

		e := NewExec(pb.ExecSpec{Command: "echo", Args: []string{"hello", "benchhub"}})
		err := e.Run(context.Background())
		assert.Nil(err)
	})
	t.Run("use context to control timeout", func(t *testing.T) {
		assert := asst.New(t)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		e := NewExec(pb.ExecSpec{Command: "sleep", Args: []string{"1"}})
		err := e.Run(ctx)
		assert.Nil(err)
		cancel()
	})
	t.Run("bash script", func(t *testing.T) {
		assert := asst.New(t)
		e := NewExec(pb.ExecSpec{Command: "testdata/echo_and_sleep.sh"})
		err := e.Run(context.Background())
		assert.Nil(err)
	})
	// TODO: test if I started a background process in shell, would it be killed when I kill the shell script
}
