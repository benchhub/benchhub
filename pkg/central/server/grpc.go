package server

import (
	"context"
	"fmt"
	igrpc "github.com/at15/go.ice/ice/transport/grpc"
	dlog "github.com/dyweb/gommon/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
	"github.com/benchhub/benchhub/pkg/central/store/meta"
	rpc "github.com/benchhub/benchhub/pkg/central/transport/grpc"
	"github.com/benchhub/benchhub/pkg/config"
)

var _ rpc.BenchHubCentralServer = (*GrpcServer)(nil)

type GrpcServer struct {
	meta         meta.Provider
	registry     *Registry
	globalConfig config.CentralServerConfig
	log          *dlog.Logger
}

func NewGrpcServer(meta meta.Provider, r *Registry) (*GrpcServer, error) {
	srv := &GrpcServer{
		meta:         meta,
		registry:     r,
		globalConfig: r.Config,
	}
	dlog.NewStructLogger(log, srv)
	return srv, nil
}

func (srv *GrpcServer) Ping(ctx context.Context, ping *pb.Ping) (*pb.Pong, error) {
	srv.log.Infof("got ping, message is %s", ping.Message)
	res := fmt.Sprintf("pong from central %s your message is %s", igrpc.Hostname(), ping.Message)
	return &pb.Pong{Message: res}, nil
}

func (srv *GrpcServer) NodeInfo(ctx context.Context, _ *pb.NodeInfoReq) (*pb.NodeInfoRes, error) {
	node, err := NodeInfo(srv.globalConfig)
	if err != nil {
		log.Warnf("failed to get central node info %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get central node info %v", err)
	}
	return &pb.NodeInfoRes{
		Node: node,
	}, nil
}

func (srv *GrpcServer) RegisterAgent(ctx context.Context, req *pb.RegisterAgentReq) (*pb.RegisterAgentRes, error) {
	remoteAddr := igrpc.RemoteAddr(ctx)
	srv.log.Infof("register agent req from %s %s %s", remoteAddr, req.Node.Host, req.Node.Id)
	req.Node.Addr.RemoteAddr = remoteAddr
	req.Node.Addr.Ip, _ = igrpc.SplitHostPort(remoteAddr)

	err := srv.meta.AddNode(req.Node.Id, pb.Node{
		Id: req.Node.Id,
		// TODO: state ... it is not passed in request, also change req.Node to req.Info ?...
		Info: req.Node,
	})
	// TODO: do something if the agent is already registered instead of just ignore it ...
	if err != nil && !pb.IsAlreadyExist(err) {
		log.Warnf("failed to add node %v", err)
		return nil, status.Errorf(codes.Internal, "failed to register node %v", err)
	}
	node, err := NodeInfo(srv.globalConfig)
	if err != nil {
		log.Warnf("failed to get central node info %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get central node info %v", err)
	}
	res := &pb.RegisterAgentRes{
		Id:      req.Node.Id,
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
