package config

import (
	"time"

	iconfig "github.com/at15/go.ice/ice/config"
)

type HeartbeatConfig struct {
	Interval time.Duration `yaml:"interval"`
}

type AgentServerConfig struct {
	Http      iconfig.HttpServerConfig `yaml:"http"`
	Grpc      iconfig.GrpcServerConfig `yaml:"grpc"`
	Node      NodeConfig               `yaml:"node"`
	Central   CentralClientConfig      `yaml:"central"`
	Heartbeat HeartbeatConfig          `yaml:"heartbeat"`
}
