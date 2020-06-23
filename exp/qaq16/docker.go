package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
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
	// --mount "type=bind,src=$(pwd)/shared,dst=/opt/shared"
	for _, mount := range cfg.Mounts {
		args = append(args, "--mount", fmt.Sprintf("type=bind,src=%s,dst=%s", mount.Src, mount.Dst))
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

	ids := bytes.Split(b, []byte{'\n'})
	var cids []string
	for _, id := range ids {
		id = bytes.TrimSpace(id)
		if len(id) == 0 {
			continue
		}
		cids = append(cids, string(id))
	}
	if len(cids) == 0 {
		log.Info("no qaq16 container to remove")
		return nil
	} else {
		log.Infof("should remove %d containers %v", len(cids), cids)
	}

	rmArgs := []string{
		"rm", "-f",
	}
	rmArgs = append(rmArgs, cids...)
	rmCmd := exec.CommandContext(ctx, "docker", rmArgs...)
	rmCmd.Stdout = os.Stdout
	rmCmd.Stderr = os.Stderr
	return errors.Wrapf(rmCmd.Run(), "error remove %d container", len(cids))
}
