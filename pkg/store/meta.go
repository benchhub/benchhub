package store

import (
	"encoding/json"
	"fmt"

	"github.com/benchhub/benchhub/bhpb"
	"github.com/dyweb/gommon/errors"
	"github.com/dyweb/gommon/util/hashutil"
)

type Meta interface {
	RegisterGoBenchmark(spec *bhpb.GoBenchmarkSpec) (*bhpb.JobRegisterResponse, error)
}

func HashGoBenchmarkSpec(spec *bhpb.GoBenchmarkSpec) (string, error) {
	// TODO: a dump way to hash entire spec, encode into bytes and hash ...
	b, err := json.Marshal(spec)
	if err != nil {
		return "", errors.Wrap(err, "error encode spec as json")
	}
	// TODO(gommon): use HashFnv64a
	return fmt.Sprintf("%x", hashutil.HashStringFnv64a(string(b))), nil
}
