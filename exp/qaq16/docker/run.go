package docker

import (
	"context"

	"github.com/benchhub/benchhub/exp/qaq16/config"
	"github.com/dyweb/gommon/errors/errortype"
)

// run docker image in foreground by shell out to docker cli

// Run blocks until context is cancelled or the container exit.
func Run(ctx context.Context, cfg config.Container) error {
	args := []string{
		"run", "--rm", "--net", "host",
	}
	for _, env := range cfg.Envs {
		args = append(args,
			"--env", env.Key, env.Value,
		)
	}
	return errortype.NewNotImplemented("Run")
}
