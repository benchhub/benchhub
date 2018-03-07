package bhpb

import (
	"testing"

	"github.com/ghodss/yaml"
	"github.com/dyweb/gommon/util/testutil"
	"time"
)

// NOTE: when no yaml tag, it will lower case it ....
type OwnerAux struct {
	S         string
	T         time.Duration
	CamelCase string
	Owner     Owner
	XXX       map[string]interface{} `yaml:",inline"`
}

// https://godoc.org/gopkg.in/yaml.v2
// Struct fields are only unmarshalled if they are exported (have an upper case first letter), and are unmarshalled using the field name lowercased as the default key
func TestOwner_YAML_Unmarshal(t *testing.T) {
	var aux OwnerAux
	testutil.ReadYAMLTo(t, "testdata/owner.yml", &aux)
	t.Log(aux.S)
	t.Log(aux.CamelCase)
	t.Log(aux.XXX)
	t.Log(aux.Owner.Name)
	t.Log(aux.Owner.Type)
}

func TestOwner_UnmarshalYAML(t *testing.T) {
	// use https://github.com/ghodss/yaml
	b := testutil.ReadFixture(t, "testdata/owner.yml")
	var aux OwnerAux
	err := yaml.Unmarshal(b, &aux)
	t.Log(aux.CamelCase)
	t.Log(aux.T)
	t.Log(aux.Owner.Name)
	t.Log(aux.Owner.Type)
	// FIXME: error unmarshaling JSON: json: cannot unmarshal string into Go struct field OwnerAux.T of type time.Duration
	t.Log(err)
}
