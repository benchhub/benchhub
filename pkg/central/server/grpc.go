package server

import (
	"context"
	"os"

	dlog "github.com/dyweb/gommon/log"
	"github.com/pkg/errors"

	pb "github.com/benchhub/benchhub/pkg/central/centralpb"
	rpc "github.com/benchhub/benchhub/pkg/central/transport/grpc"
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

func (srv *GrpcServer) Ping(ctx context.Context, ping *pb.Ping) (*pb.Pong, error) {
	if host, err := os.Hostname(); err != nil {
		return &pb.Pong{Name: "unknown"}, errors.Wrap(err, "can't get hostname")
	} else {
		return &pb.Pong{Name: host}, nil
	}
}

func (srv *GrpcServer) RegisterAgent(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterRes, error) {
	// TODO: https://groups.google.com/forum/#!topic/grpc-io/UodEY4N78Sk
	// tell the agent what its address in central's perspective,
	//peer, err := peer.FromContext(ctx)
	//peer.Addr
	return nil, nil
}
