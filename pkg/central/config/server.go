package config

import (
	"github.com/at15/go.ice/ice/config"
)

type ServerConfig struct {
	Grpc config.GrpcServerConfig `yaml:"grpc"`
}
