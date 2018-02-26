package config

import (
	iconfig "github.com/at15/go.ice/ice/config"
	cconfig "github.com/benchhub/benchhub/pkg/common/config"
)

type CentralConfig struct {
	Addr string `yaml:"addr"`
}

type ServerConfig struct {
	Http    iconfig.HttpServerConfig `yaml:"http"`
	Grpc    iconfig.GrpcServerConfig `yaml:"grpc"`
	Central CentralConfig            `yaml:"central"`
	Node    cconfig.NodeConfig       `yaml:"node"`
}
