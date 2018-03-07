package runner

import (
	"context"
	"testing"

	"github.com/dyweb/gommon/util/testutil"
	asst "github.com/stretchr/testify/assert"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

func TestDocker_Pull(t *testing.T) {
	if testing.Short() {
		t.Skipf("skip runner test in short tests")
	}

	testutil.RunIf(t, testutil.IsTravis())

	t.Run("pull", func(t *testing.T) {
		assert := asst.New(t)
		d, err := NewDocker(pb.DockerSpec{
			Image:  "influxdb:1.3.9",
			Action: pb.DockerAction_PULL,
		}, nil)
		assert.Nil(err)
		err = d.Run(context.Background())
		assert.Nil(err)
	})
	t.Run("pull not exist", func(t *testing.T) {
		assert := asst.New(t)
		d, err := NewDocker(pb.DockerSpec{
			Image:  "xephonk:2.0",
			Action: pb.DockerAction_PULL,
		}, nil)
		assert.Nil(err)
		err = d.Run(context.Background())
		if err != nil {
			t.Log(err.Error())
		}
		assert.NotNil(err)
	})
	t.Run("start", func(t *testing.T) {
		assert := asst.New(t)
		d, err := NewDocker(pb.DockerSpec{
			Image:  "influxdb:1.3.9",
			Action: pb.DockerAction_RUN,
		}, nil)
		assert.Nil(err)
		err = d.Run(context.Background())
		if err != nil {
			t.Log(err.Error())
		}
		assert.Nil(err)
	})
}
