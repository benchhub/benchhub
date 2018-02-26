package server

import (
	"context"
	"fmt"
	"os"

	dlog "github.com/dyweb/gommon/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/benchhub/benchhub/pkg/central/centralpb"
	"github.com/benchhub/benchhub/pkg/central/store/meta"
	rpc "github.com/benchhub/benchhub/pkg/central/transport/grpc"
	pbc "github.com/benchhub/benchhub/pkg/common/commonpb"
	"github.com/benchhub/benchhub/pkg/common/nodeutil"
)

var _ rpc.BenchHubCentralServer = (*GrpcServer)(nil)

type GrpcServer struct {
	meta meta.Provider
	log  *dlog.Logger
}

func NewGrpcServer(meta meta.Provider) (*GrpcServer, error) {
	srv := &GrpcServer{
		meta: meta,
	}
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
	res := fmt.Sprintf("pong from central %s your message is %s", hostname(), ping.Message)
	return &pbc.Pong{Message: res}, nil
}

func (srv *GrpcServer) RegisterAgent(ctx context.Context, req *pb.RegisterAgentReq) (*pb.RegisterAgentRes, error) {
	// TODO:
	// - check if the node is already registered
	// - assign it id
	// - return information about itself
	srv.log.Infof("register agent req from %s", req.Node.Host)
	err := srv.meta.AddNode(req.Node.Uid, req.Node)
	if err != nil {
		log.Warnf("failed to add node %v", err)
		// TODO: already exists may not be the only cause .... though for in memory, it should be ...
		return nil, status.Errorf(codes.AlreadyExists, "failed to add node %v", err)
	}
	central, err := nodeutil.GetNode()
	// TODO: update bindAddr, ip, port, etc.
	if err != nil {
		log.Warnf("failed to get central node info %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get central node info %v", err)
	}
	res := &pb.RegisterAgentRes{
		Id:      req.Node.Uid,
		Node:    req.Node,
		Central: *central,
	}
	return res, nil
}

func (srv *GrpcServer) AgentHeartbeat(ctx context.Context, req *pb.AgentHeartbeatReq) (*pb.AgentHeartbeatRes, error) {
	return nil, status.Error(codes.Unimplemented, "heartbeat is under construction")
}

func hostname() string {
	if host, err := os.Hostname(); err != nil {
		log.Warnf("can't get hostname %v", err)
		return "unknown"
	} else {
		return host
	}
}
