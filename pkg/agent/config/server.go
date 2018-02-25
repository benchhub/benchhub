package config

import (
	"github.com/at15/go.ice/ice/config"
)

type CentralConfig struct {
	Addr string `yaml:"addr"`
}

type ServerConfig struct {
	Http    config.HttpServerConfig `yaml:"http"`
	Grpc    config.GrpcServerConfig `yaml:"grpc"`
	Central CentralConfig           `yaml:"central"`
}
