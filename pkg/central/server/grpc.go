package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"

	igrpc "github.com/at15/go.ice/ice/transport/grpc"
	dlog "github.com/dyweb/gommon/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/benchhub/benchhub/pkg/central/centralpb"
	"github.com/benchhub/benchhub/pkg/central/config"
	"github.com/benchhub/benchhub/pkg/central/store/meta"
	rpc "github.com/benchhub/benchhub/pkg/central/transport/grpc"
	pbc "github.com/benchhub/benchhub/pkg/common/commonpb"
	"github.com/benchhub/benchhub/pkg/common/nodeutil"
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
	res := fmt.Sprintf("pong from central %s your message is %s", hostname(), ping.Message)
	return &pbc.Pong{Message: res}, nil
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

	node, err := nodeutil.GetNode()
	if err != nil {
		log.Warnf("failed to get central node info %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get central node info %v", err)
	}
	node.BindAdrr = srv.globalConfig.Grpc.Addr
	node.BindIp, node.BindPort = splitHostPort(node.BindAdrr)
	node.Provider = pbc.NodeProvider{
		Name:     srv.globalConfig.Node.Provider.Name,
		Region:   srv.globalConfig.Node.Provider.Region,
		Instance: srv.globalConfig.Node.Provider.Instance,
	}
	res := &pb.RegisterAgentRes{
		Id:      req.Node.Uid,
		Node:    req.Node,
		Central: *node,
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

// FIXME: exact duplicated code in central and agent, this should go to go.ice
func splitHostPort(addr string) (string, int64) {
	_, ps, err := net.SplitHostPort(addr)
	if err != nil {
		log.Warnf("failed to split host port %s %v", addr, err)
		return "", 0
	}
	p, err := strconv.Atoi(ps)
	if err != nil {
		log.Warnf("failed to convert port number %s to int %v", ps, err)
		return ps, int64(p)
	}
	return ps, int64(p)
}
