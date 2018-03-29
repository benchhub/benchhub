package config

import (
	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

type NodeConfig struct {
	// Role is preferred role of this node, should be set based on instance type
	Role     pb.Role         `yaml:"role"`
	Provider pb.NodeProvider `yaml:"provider"`
}
