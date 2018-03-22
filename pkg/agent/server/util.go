package server

import (
	"github.com/benchhub/benchhub/pkg/agent/config"
	pb "github.com/benchhub/benchhub/pkg/bhpb"
	"github.com/benchhub/benchhub/pkg/common/nodeutil"
)

// FIXME: this is dup in both agent and central
func NodeInfo(cfg config.ServerConfig) (*pb.NodeInfo, error) {
	node, err := nodeutil.GetNodeInfo(cfg.Node)
	node.Addr = pb.Addr{
		BindAddr: cfg.Grpc.Addr,
	}
	if err != nil {
		return node, err
	}
	return node, nil
}
