package bhpb

import (
	"testing"
	"time"
	"encoding/json"

	"github.com/ghodss/yaml"
	"github.com/dyweb/gommon/util/testutil"
	asst "github.com/stretchr/testify/assert"
)

// NOTE: when no yaml tag, it will lower case it ....
type OwnerWrapper struct {
	S         string
	T         time.Duration
	CamelCase string
	Owner     Owner
	XXX       map[string]interface{} `yaml:",inline"`
}

// how to handle time.Duration when using ghodss/yaml, it is supported by go-yaml, but not json
// http://choly.ca/post/go-json-marshalling/
func (o *OwnerWrapper) UnmarshalJSON(data []byte) error {
	type Alias OwnerWrapper
	aux := &struct {
		T string
		*Alias
	}{
		Alias: (*Alias)(o),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if d, err := time.ParseDuration(aux.T); err != nil {
		return err
	} else {
		o.T = d
	}
	return nil
}

// using https://godoc.org/gopkg.in/yaml.v2
// Struct fields are only unmarshalled if they are exported (have an upper case first letter), and are unmarshalled using the field name lowercased as the default key
func TestOwner_YAML_Unmarshal(t *testing.T) {
	var wrapper OwnerWrapper
	testutil.ReadYAMLTo(t, "testdata/owner.yml", &wrapper)
	t.Log(wrapper.S)
	t.Log(wrapper.CamelCase)
	t.Log(wrapper.XXX)
	t.Log(wrapper.Owner.Name)
	t.Log(wrapper.Owner.Type)
}

func TestOwner_UnmarshalYAML(t *testing.T) {
	assert := asst.New(t)

	// use https://github.com/ghodss/yaml
	b := testutil.ReadFixture(t, "testdata/owner.yml")
	var wrapper OwnerWrapper
	err := yaml.Unmarshal(b, &wrapper)
	assert.Nil(err)
	assert.Equal("camelallsmall", wrapper.CamelCase)
	// FIXED: error unmarshaling JSON: json: cannot unmarshal string into Go struct field OwnerAux.T of type time.Duration
	assert.Equal(10*time.Second, wrapper.T)
	assert.Equal("at15", wrapper.Owner.Name)
	assert.Equal(OwnerType_ORG, wrapper.Owner.Type)
}
