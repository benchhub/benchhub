package config

import (
	"github.com/at15/go.ice/ice/config"
)

type ServerConfig struct {
	Http config.HttpServerConfig `yaml:"http"`
	Grpc config.GrpcServerConfig `yaml:"grpc"`

	// TODO: it need to know how to reach out to central
}
