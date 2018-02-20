package server

import (
	"context"
	"os"

	"github.com/pkg/errors"

	pb "github.com/benchhub/benchhub/pkg/central/centralpb"
	rpc "github.com/benchhub/benchhub/pkg/central/transport/grpc"
)

var _ rpc.BenchHubCentralServer = (*GrpcServer)(nil)

type GrpcServer struct {
}

func NewGrpcServer() (*GrpcServer, error) {
	return &GrpcServer{}, nil
}

func (srv *GrpcServer) Ping(ctx context.Context, ping *pb.Ping) (*pb.Pong, error) {
	if host, err := os.Hostname(); err != nil {
		return &pb.Pong{Name: "unknown"}, errors.Wrap(err, "can't get hostname")
	} else {
		return &pb.Pong{Name: host}, nil
	}
}
