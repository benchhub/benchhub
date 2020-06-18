package main

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/benchhub/benchhub/exp/qaq16/config"
	"github.com/dyweb/gommon/errors"
	"github.com/dyweb/gommon/errors/errortype"
)

func RunScore(ctx context.Context, cfg config.Score) error {
	// TODO: should use shell split args etc.
	cmdCfg := cfg.Command
	args := strings.Split(cmdCfg.Shell, " ")
	bin := args[0]
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	fullBin := filepath.Join(pwd, bin)
	log.Infof("bin %s timeout %s", fullBin, cfg.Timeout)

	timeout, err := time.ParseDuration(cfg.Timeout)
	if err != nil {
		return errors.Wrap(err, "invalid timeout for command")
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, fullBin, args[1:]...)
	cmd.Dir = cmdCfg.Dir

	// TODO: stream log to file, and optionally stdout/stderr
	return errortype.NewNotImplemented("RunCommand")
}
