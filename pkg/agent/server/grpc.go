package server

import (
	"context"
	"fmt"
	"os"

	dlog "github.com/dyweb/gommon/log"
	"github.com/pkg/errors"

	rpc "github.com/benchhub/benchhub/pkg/agent/transport/grpc"
	pbc "github.com/benchhub/benchhub/pkg/common/commonpb"
)

var _ rpc.BenchHubAgentServer = (*GrpcServer)(nil)

type GrpcServer struct {
	log *dlog.Logger
}

// TODO: it might need registry
func NewGrpcServer() (*GrpcServer, error) {
	srv := &GrpcServer{}
	dlog.NewStructLogger(log, srv)
	return srv, nil
}

// TODO: get peer information
// TODO: https://groups.google.com/forum/#!topic/grpc-io/UodEY4N78Sk
// tell the agent what its address in central's perspective,
//peer, err := peer.FromContext(ctx)
//peer.Addr

func (srv *GrpcServer) Ping(ctx context.Context, ping *pbc.Ping) (*pbc.Pong, error) {
	srv.log.Infof("got ping, message is %s", ping.Message)
	if host, err := os.Hostname(); err != nil {
		return &pbc.Pong{Message: "pong from unknown"}, errors.Wrap(err, "can't get hostname")
	} else {
		res := fmt.Sprintf("pong from agent %s your message is %s", host, ping.Message)
		return &pbc.Pong{Message: res}, nil
	}
}
