package config

import (
	iconfig "github.com/at15/go.ice/ice/config"
	cconfig "github.com/benchhub/benchhub/pkg/config"
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
	Node      cconfig.NodeConfig       `yaml:"node"`
	Central   CentralConfig            `yaml:"central"`
	Heartbeat HeartbeatConfig          `yaml:"heartbeat"`
}
