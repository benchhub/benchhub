package main

import (
	"context"
	"os/exec"
	"strings"

	"github.com/benchhub/benchhub/exp/qaq16/config"
	"github.com/dyweb/gommon/errors"
)

// RunContainer blocks until context is cancelled or the container exit.
func RunContainer(ctx context.Context, cfg config.Container, run ExecContext) error {
	args := []string{
		"run", "--rm", "--net", "host",
		"--label", "qaq16=1", // assign label so we can kill all of them
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

func RmContainers(ctx context.Context) error {
	psArgs := []string{
		"ps", "-f", "label=qaq16=1", "-q",
	}
	psCmd := exec.CommandContext(ctx, "docker", psArgs...)
	b, err := psCmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, "error find qaq16 containers")
	}
	log.Infof("%s", b)
	// TODO: extract ids and docker rm, split by new line
	return nil
}
