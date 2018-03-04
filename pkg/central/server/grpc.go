package server

import (
	"context"
	"fmt"
	igrpc "github.com/at15/go.ice/ice/transport/grpc"
	dlog "github.com/dyweb/gommon/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/benchhub/benchhub/pkg/central/centralpb"
	"github.com/benchhub/benchhub/pkg/central/config"
	"github.com/benchhub/benchhub/pkg/central/store/meta"
	rpc "github.com/benchhub/benchhub/pkg/central/transport/grpc"
	pbc "github.com/benchhub/benchhub/pkg/common/commonpb"
)

var _ rpc.BenchHubCentralServer = (*GrpcServer)(nil)

type GrpcServer struct {
	meta         meta.Provider
	globalConfig config.ServerConfig
	log          *dlog.Logger
}

func NewGrpcServer(meta meta.Provider, cfg config.ServerConfig) (*GrpcServer, error) {
	srv := &GrpcServer{
		meta:         meta,
		globalConfig: cfg,
	}
	dlog.NewStructLogger(log, srv)
	return srv, nil
}

func (srv *GrpcServer) Ping(ctx context.Context, ping *pbc.Ping) (*pbc.Pong, error) {
	srv.log.Infof("got ping, message is %s", ping.Message)
	res := fmt.Sprintf("pong from central %s your message is %s", igrpc.Hostname(), ping.Message)
	return &pbc.Pong{Message: res}, nil
}

func (srv *GrpcServer) NodeInfo(ctx context.Context, _ *pbc.NodeInfoReq) (*pbc.NodeInfoRes, error) {
	node, err := Node(srv.globalConfig)
	if err != nil {
		log.Warnf("failed to get central node info %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get central node info %v", err)
	}
	return &pbc.NodeInfoRes{
		Node: node,
	}, nil
}

func (srv *GrpcServer) RegisterAgent(ctx context.Context, req *pb.RegisterAgentReq) (*pb.RegisterAgentRes, error) {
	remoteAddr := igrpc.RemoteAddr(ctx)
	srv.log.Infof("register agent req from %s %s", remoteAddr, req.Node.Host)
	req.Node.RemoteAddr = remoteAddr

	err := srv.meta.AddNode(req.Node.Uid, req.Node)
	if err != nil {
		log.Warnf("failed to add node %v", err)
		// TODO: already exists may not be the only cause .... though for in memory, it should be ...
		return nil, status.Errorf(codes.AlreadyExists, "failed to add node %v", err)
	}
	node, err := Node(srv.globalConfig)
	if err != nil {
		log.Warnf("failed to get central node info %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get central node info %v", err)
	}
	res := &pb.RegisterAgentRes{
		Id:      req.Node.Uid,
		Node:    req.Node,
		Central: *node,
	}
	return res, nil
}

func (srv *GrpcServer) AgentHeartbeat(ctx context.Context, req *pb.AgentHeartbeatReq) (*pb.AgentHeartbeatRes, error) {
	if err := srv.meta.UpdateNodeStatus(req.Id, req.Status); err != nil {
		log.Warnf("failed to update status for %s %v", req.Id, err)
		return nil, status.Errorf(codes.Internal, "failed to update status for %s %v", req.Id, err)
	}
	return &pb.AgentHeartbeatRes{}, nil
}

func (srv *GrpcServer) ListAgent(ctx context.Context, req *pb.ListAgentReq) (*pb.ListAgentRes, error) {
	node, err := srv.meta.ListNodes()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list nodes %s", err.Error())
	}
	return &pb.ListAgentRes{
		Agents: node,
	}, nil
}
