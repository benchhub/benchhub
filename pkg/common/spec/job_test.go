package spec

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	tu "github.com/dyweb/gommon/util/testutil"
	asst "github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestJob_Parse(t *testing.T) {

	files := []string{
		//"pingpong.yml",
		"xephonb-kairosdb.yml",
	}

	for _, f := range files {
		t.Run(f, func(t *testing.T) {
			assert := asst.New(t)
			data := tu.ReadFixture(t, f)
			var job Job
			err := yaml.Unmarshal(data, &job)

			assert.Nil(err)
			err = job.Validate()
			if err != nil {
				t.Log(err.Error())
			}
			assert.Nil(err)
			if tu.Dump().B() {
				spew.Dump(job)
			}
		})
	}

}
