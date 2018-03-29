package server

import (
	pb "github.com/benchhub/benchhub/pkg/bhpb"
	"github.com/benchhub/benchhub/pkg/config"
	"github.com/benchhub/benchhub/pkg/util/nodeutil"
)

// FIXME: this is dup in both agent and central
func NodeInfo(cfg config.AgentServerConfig) (*pb.NodeInfo, error) {
	node, err := nodeutil.GetNodeInfo(cfg.Node)
	node.Addr = pb.Addr{
		BindAddr: cfg.Grpc.Addr,
	}
	if err != nil {
		return node, err
	}
	return node, nil
}
