package server

import (
	"context"
	"fmt"
	"os"

	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"

	pb "github.com/benchhub/benchhub/pkg/central/centralpb"
	rpc "github.com/benchhub/benchhub/pkg/central/transport/grpc"
	pbc "github.com/benchhub/benchhub/pkg/common/commonpb"
)

var _ rpc.BenchHubCentralServer = (*GrpcServer)(nil)

type GrpcServer struct {
	log *dlog.Logger
}

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
		return &pbc.Pong{Message: "pong from agent unknown"}, errors.Wrap(err, "can't get hostname")
	} else {
		res := fmt.Sprintf("pong from central %s your message is %s", host, ping.Message)
		return &pbc.Pong{Message: res}, nil
	}
}

func (srv *GrpcServer) RegisterAgent(ctx context.Context, req *pb.RegisterAgentReq) (*pb.RegisterAgentRes, error) {
	// TODO:
	// - check if the node is already registered
	// - assign it id
	// - return information about itself
	return nil, nil
}

func (srv *GrpcServer) AgentHeartbeat(ctx context.Context, req *pb.AgentHeartbeatReq) (*pb.AgentHeartbeatRes, error) {
	return nil, nil
}
