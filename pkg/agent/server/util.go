package server

import (
	"github.com/benchhub/benchhub/pkg/agent/config"
	pb "github.com/benchhub/benchhub/pkg/bhpb"
	"github.com/benchhub/benchhub/pkg/common/nodeutil"
)

// FIXME: this is dup in both agent and central
func NodeInfo(cfg config.ServerConfig) (*pb.NodeInfo, error) {
	node, err := nodeutil.GetNodeInfo()
	node.Addr = pb.Addr{
		BindAddr: cfg.Grpc.Addr,
	}
	node.Provider = pb.NodeProvider{
		Name:     cfg.Node.Provider.Name,
		Region:   cfg.Node.Provider.Region,
		Instance: cfg.Node.Provider.Instance,
	}
	node.Role = pb.Role(pb.Role_value[cfg.Node.Role])
	if err != nil {
		return node, err
	}
	return node, nil
}
