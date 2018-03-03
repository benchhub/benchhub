package runner

import (
	"testing"
	"context"

	asst "github.com/stretchr/testify/assert"

	"github.com/benchhub/benchhub/pkg/common/spec"
)

func TestDocker_Run(t *testing.T) {
	if testing.Short() {
		t.Skipf("skip runner test in short tests")
	}
	t.Run("pull", func(t *testing.T) {
		assert := asst.New(t)
		d, err := NewDocker(spec.Docker{
			Image:  "influxdb:1.3.9",
			Action: spec.DockerPull,
		}, nil)
		assert.Nil(err)
		err = d.Run(context.Background())
		assert.Nil(err)
	})
	t.Run("pull not exist", func(t *testing.T) {
		assert := asst.New(t)
		d, err := NewDocker(spec.Docker{
			Image:  "xephonk:2.0",
			Action: spec.DockerPull,
		}, nil)
		assert.Nil(err)
		err = d.Run(context.Background())
		if err != nil {
			t.Log(err.Error())
		}
		assert.NotNil(err)
	})
}
