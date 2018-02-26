package config

import (
	iconfig "github.com/at15/go.ice/ice/config"
)

type MetaConfig struct {
	Provider string `yaml:"provider"`
}

type ServerConfig struct {
	Http iconfig.HttpServerConfig `yaml:"http"`
	Grpc iconfig.GrpcServerConfig `yaml:"grpc"`
	Meta MetaConfig               `yaml:"meta"`
}
