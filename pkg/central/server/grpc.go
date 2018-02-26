package server

import (
	"context"
	"fmt"
	"os"

	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"

	pb "github.com/benchhub/benchhub/pkg/central/centralpb"
	"github.com/benchhub/benchhub/pkg/central/store/meta"
	rpc "github.com/benchhub/benchhub/pkg/central/transport/grpc"
	pbc "github.com/benchhub/benchhub/pkg/common/commonpb"
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
	srv.log.Infof("got register from %s", req.Node.Host)
	// TODO: use grpc/codes and grpc.Errorf https://groups.google.com/forum/#!topic/golang-nuts/NZX1sOYosRs
	// https://godoc.org/google.golang.org/grpc/status
	return nil, errors.New("not implemented")
}

func (srv *GrpcServer) AgentHeartbeat(ctx context.Context, req *pb.AgentHeartbeatReq) (*pb.AgentHeartbeatRes, error) {
	return nil, errors.New("not implemented")
}

func hostname() string {
	if host, err := os.Hostname(); err != nil {
		log.Warnf("can't get hostname %v", err)
		return "unknown"
	} else {
		return host
	}
}
