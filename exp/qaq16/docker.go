package main

import (
	"context"
	"os/exec"
	"strings"

	"github.com/benchhub/benchhub/exp/qaq16/config"
	"github.com/dyweb/gommon/errors"
)

// RunContainer blocks until context is cancelled or the container exit.
// TODO: assign label and allow kill all the containers ...
func RunContainer(ctx context.Context, cfg config.Container, run ExecContext) error {
	args := []string{
		"run", "--rm", "--net", "host",
	}
	for _, env := range cfg.Envs {
		args = append(args,
			"--env", env.Key+"="+env.Value,
		)
	}
	args = append(args, cfg.Image)
	log.Infof("docker %s", strings.Join(args, " "))
	cmd := exec.CommandContext(ctx, "docker", args...)
	return errors.Wrap(RunCommand(cmd, run.log), cfg.Name)
}
