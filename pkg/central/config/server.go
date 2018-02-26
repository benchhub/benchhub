package config

import (
	"github.com/at15/go.ice/ice/config"
)

type MetaConfig struct {
	Provider string `yaml:"provider"`
}

type ServerConfig struct {
	Http config.HttpServerConfig `yaml:"http"`
	Grpc config.GrpcServerConfig `yaml:"grpc"`
	Meta MetaConfig              `yaml:"meta"`
}
