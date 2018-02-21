package server

import (
	"context"
	"os"

	dlog "github.com/dyweb/gommon/log"
	"github.com/pkg/errors"

	pb "github.com/benchhub/benchhub/pkg/agent/agentpb"
	rpc "github.com/benchhub/benchhub/pkg/agent/transport/grpc"
)

var _ rpc.BenchHubAgentServer = (*GrpcServer)(nil)

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
