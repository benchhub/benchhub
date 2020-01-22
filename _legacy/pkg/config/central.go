package config

import (
	iconfig "github.com/at15/go.ice/ice/config"
	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

type CentralClientConfig struct {
	Addr string `yaml:"addr"`
}

type CentralServerConfig struct {
	Http iconfig.HttpServerConfig `yaml:"http"`
	Grpc iconfig.GrpcServerConfig `yaml:"grpc"`
	Node pb.NodeConfig            `yaml:"node"`
	Meta MetaStoreConfig          `yaml:"meta"`
	Job  JobConfig                `yaml:"job"`
}
