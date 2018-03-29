package config

import (
	"time"

	iconfig "github.com/at15/go.ice/ice/config"
	cconfig "github.com/benchhub/benchhub/pkg/config"
)

type MetaConfig struct {
	Provider string `yaml:"provider"`
}

type JobConfig struct {
	PollInterval time.Duration `yaml:"pollInterval"`
}

type ServerConfig struct {
	Http iconfig.HttpServerConfig `yaml:"http"`
	Grpc iconfig.GrpcServerConfig `yaml:"grpc"`
	Node cconfig.NodeConfig       `yaml:"node"`
	Meta MetaConfig               `yaml:"meta"`
	Job  JobConfig                `yaml:"job"`
}
