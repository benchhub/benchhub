package config

import (
	"time"

	iconfig "github.com/at15/go.ice/ice/config"
	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

type HeartbeatConfig struct {
	Interval time.Duration `yaml:"interval"`
}

type AgentServerConfig struct {
	Http      iconfig.HttpServerConfig `yaml:"http"`
	Grpc      iconfig.GrpcServerConfig `yaml:"grpc"`
	Node      pb.NodeConfig            `yaml:"node"`
	Central   CentralClientConfig      `yaml:"central"`
	Heartbeat HeartbeatConfig          `yaml:"heartbeat"`
}
