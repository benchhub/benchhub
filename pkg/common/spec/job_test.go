package spec

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/dyweb/gommon/util/testutil"
	asst "github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestJob_Parse(t *testing.T) {
	assert := asst.New(t)

	data := testutil.ReadFixture(t, "xephonb-kairosdb.yml")
	var job Job
	err := yaml.Unmarshal(data, &job)

	assert.Nil(err)
	err = job.Validate()
	if err != nil {
		t.Log(err.Error())
	}
	assert.Nil(err)
	// TODO: need to handle config of task, is using interface ...
	spew.Dump(job)
}
