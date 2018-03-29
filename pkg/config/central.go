package config

import (
	iconfig "github.com/at15/go.ice/ice/config"
)

type CentralClientConfig struct {
	Addr string `yaml:"addr"`
}

type CentralServerConfig struct {
	Http iconfig.HttpServerConfig `yaml:"http"`
	Grpc iconfig.GrpcServerConfig `yaml:"grpc"`
	Node NodeConfig               `yaml:"node"`
	Meta MetaStoreConfig          `yaml:"meta"`
	Job  JobConfig                `yaml:"job"`
}
