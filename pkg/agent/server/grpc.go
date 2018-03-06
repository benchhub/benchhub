package server

import (
	"context"
	"fmt"

	igrpc "github.com/at15/go.ice/ice/transport/grpc"
	dlog "github.com/dyweb/gommon/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/benchhub/benchhub/pkg/agent/config"
	rpc "github.com/benchhub/benchhub/pkg/agent/transport/grpc"
	pbc "github.com/benchhub/benchhub/pkg/bhpb"
)

var _ rpc.BenchHubAgentServer = (*GrpcServer)(nil)

type GrpcServer struct {
	registry     *Registry
	globalConfig config.ServerConfig
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

func (srv *GrpcServer) Ping(ctx context.Context, ping *pbc.Ping) (*pbc.Pong, error) {
	srv.log.Infof("got ping, message is %s", ping.Message)
	res := fmt.Sprintf("pong from agent %s your message is %s", igrpc.Hostname(), ping.Message)
	return &pbc.Pong{Message: res}, nil
}

func (srv *GrpcServer) NodeInfo(ctx context.Context, _ *pbc.NodeInfoReq) (*pbc.NodeInfoRes, error) {
	node, err := NodeInfo(srv.globalConfig)
	if err != nil {
		log.Warnf("failed to get central node info %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get central node info %v", err)
	}
	return &pbc.NodeInfoRes{
		Node: node,
	}, nil
}
