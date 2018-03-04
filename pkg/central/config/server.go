package config

import (
	iconfig "github.com/at15/go.ice/ice/config"
	cconfig "github.com/benchhub/benchhub/pkg/common/config"
)

type MetaConfig struct {
	Provider string `yaml:"provider"`
}

type ServerConfig struct {
	Http iconfig.HttpServerConfig `yaml:"http"`
	Grpc iconfig.GrpcServerConfig `yaml:"grpc"`
	Meta MetaConfig               `yaml:"meta"`
	Node cconfig.NodeConfig       `yaml:"node"`
}
