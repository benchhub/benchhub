package server

import (
	"github.com/benchhub/benchhub/pkg/central/config"
	"github.com/benchhub/benchhub/pkg/central/store/meta"
)

type Registry struct {
	Meta   meta.Provider
	Config config.ServerConfig
}
