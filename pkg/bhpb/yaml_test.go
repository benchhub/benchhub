package bhpb

import (
	"testing"
	"github.com/dyweb/gommon/util/testutil"
)

func TestOwner_YAML_Unmarshal(t *testing.T) {
	owner := Owner{}
	testutil.ReadYAMLTo(t, "testdata/owner.yml", &owner)
	t.Log(owner.Name)
	t.Log(owner.Type)
}