package server

import (
	igrpc "github.com/at15/go.ice/ice/transport/grpc"

	"github.com/benchhub/benchhub/pkg/agent/config"
	pbc "github.com/benchhub/benchhub/pkg/common/commonpb"
	"github.com/benchhub/benchhub/pkg/common/nodeutil"
)

// TODO: this is dup in both agent and central
func Node(cfg config.ServerConfig) (*pbc.Node, error) {
	node, err := nodeutil.GetNode()
	node.BindAdrr = cfg.Grpc.Addr
	node.BindIp, node.BindPort = igrpc.SplitHostPort(node.BindAdrr)
	node.Provider = pbc.NodeProvider{
		Name:     cfg.Node.Provider.Name,
		Region:   cfg.Node.Provider.Region,
		Instance: cfg.Node.Provider.Instance,
	}
	node.Role = pbc.NodeRole{
		// TODO: better way to use enumerate
		Preferred: pbc.Role(pbc.Role_value[cfg.Node.Role]),
		// TODO: need to know previous and current role ....
	}
	if err != nil {
		return node, err
	}
	return node, nil
}
