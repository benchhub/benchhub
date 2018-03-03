package config

import (
	iconfig "github.com/at15/go.ice/ice/config"
	cconfig "github.com/benchhub/benchhub/pkg/common/config"
)

type CentralConfig struct {
	Addr string `yaml:"addr"`
}

type HeartbeatConfig struct {
	Interval string `yaml:"interval"` // TODO: this requires time.ParseDuration
}

type ServerConfig struct {
	Http      iconfig.HttpServerConfig `yaml:"http"`
	Grpc      iconfig.GrpcServerConfig `yaml:"grpc"`
	Central   CentralConfig            `yaml:"central"`
	Node      cconfig.NodeConfig       `yaml:"node"`
	Heartbeat HeartbeatConfig          `yaml:"heartbeat"`
}
