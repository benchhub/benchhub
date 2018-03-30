package server

import (
	"context"
	"fmt"

	igrpc "github.com/at15/go.ice/ice/transport/grpc"
	rpc "github.com/benchhub/benchhub/pkg/agent/transport/grpc"
	pb "github.com/benchhub/benchhub/pkg/bhpb"
	"github.com/benchhub/benchhub/pkg/config"
	dlog "github.com/dyweb/gommon/log"
)

var _ rpc.BenchHubAgentServer = (*GrpcServer)(nil)

type GrpcServer struct {
	registry     *Registry
	globalConfig config.AgentServerConfig
	log          *dlog.Logger
}

func NewGrpcServer(r *Registry) (*GrpcServer, error) {
	srv := &GrpcServer{
		registry:     r,
		globalConfig: r.Config,
	}
	dlog.NewStructLogger(log, srv)
	return srv, nil
}

func (srv *GrpcServer) Ping(ctx context.Context, ping *pb.Ping) (*pb.Pong, error) {
	srv.log.Infof("got ping, message is %s", ping.Message)
	res := fmt.Sprintf("pong from agent %s your message is %s", igrpc.Hostname(), ping.Message)
	return &pb.Pong{Message: res}, nil
}

func (srv *GrpcServer) NodeInfo(ctx context.Context, _ *pb.NodeInfoReq) (*pb.NodeInfoRes, error) {
	return &pb.NodeInfoRes{
		Node: srv.registry.NodeInfo(),
	}, nil
}
