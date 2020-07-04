package main

import (
	"bytes"
	"context"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/benchhub/benchhub/exp/qaq16/config"
	"github.com/dyweb/gommon/errors"
)

type ExecContext struct {
	log string
}

func RunScore(ctx context.Context, cfg config.Score, run ExecContext) error {
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
	cmd.Dir = filepath.Join(pwd, cmdCfg.Dir)

	return RunCommand(cmd, run.log)
}

func RunCommand(cmd *exec.Cmd, logPath string) error {
	logFile, err := os.Create(logPath)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	stdout := io.MultiWriter(os.Stdout, logFile, &buf)
	stderr := io.MultiWriter(os.Stderr, logFile, &buf)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	return cmd.Run()
}
