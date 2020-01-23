package store

import (
	"context"
	"fmt"

	"github.com/benchhub/benchhub/bhpb"
	"github.com/dyweb/gommon/errors"
	"github.com/dyweb/gommon/util/hashutil"
	"github.com/gogo/protobuf/proto"
)

type Meta interface {
	RegisterGoBenchmark(ctx context.Context, spec *bhpb.GoBenchmarkSpec) (*bhpb.JobRegisterResponse, error)
}

func HashGoBenchmarkSpec(spec *bhpb.GoBenchmarkSpec) (string, error) {
	// TODO: a dump way to hash entire spec, encode into bytes and hash ...
	b, err := proto.Marshal(spec)
	if err != nil {
		return "", errors.Wrap(err, "error encode spec as json")
	}
	// TODO(gommon): use HashFnv64a
	return fmt.Sprintf("%x", hashutil.HashStringFnv64a(string(b))), nil
}
