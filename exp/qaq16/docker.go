package main

import (
	"context"
	"strings"

	"github.com/benchhub/benchhub/exp/qaq16/config"
	"github.com/dyweb/gommon/errors/errortype"
)

// RunContainer blocks until context is cancelled or the container exit.
func RunContainer(ctx context.Context, cfg config.Container) error {
	args := []string{
		"run", "--rm", "--net", "host",
	}
	for _, env := range cfg.Envs {
		args = append(args,
			"--env", env.Key, env.Value,
		)
	}
	log.Infof("docker %s", strings.Join(args, " "))
	return errortype.NewNotImplemented("RunContainer")
}
