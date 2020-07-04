package store

import (
	"context"
	"fmt"

	"github.com/benchhub/benchhub/bhpb"
	"github.com/dyweb/gommon/errors"
	"github.com/dyweb/gommon/util/hashutil"
	"github.com/gogo/protobuf/proto"
)

// TODO: the store interface is doing business logic, need to split those out
type Meta interface {
	GoBenchmarkRegister(ctx context.Context, spec *bhpb.GoBenchmarkSpec) (*bhpb.JobRegisterResponse, error)
	GoBenchmarkReportResult(ctx context.Context, result *bhpb.GoBenchmarkReportResultRequest) (*bhpb.ResultReportResponse, error)
}

// TODO: might use sha1 to reduce chance of collision?
func HashGoBenchmarkSpec(spec *bhpb.GoBenchmarkSpec) (string, error) {
	// TODO: a dump way to hash entire spec, encode into bytes and hash ...
	b, err := proto.Marshal(spec)
	if err != nil {
		return "", errors.Wrap(err, "error encode spec as json")
	}
	// TODO(gommon): use HashFnv64a
	return fmt.Sprintf("%x", hashutil.HashStringFnv64a(string(b))), nil
}
