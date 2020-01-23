package main

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/benchhub/benchhub/bhpb"
	"github.com/benchhub/benchhub/pkg/gobench"
	"github.com/dyweb/gommon/errors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// run.go allocate a job id, run the benchmark and report result

const DefaultConfig = "benchhub.yml"

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run benchmark",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := mustDefaultClient()
		var spec bhpb.GoBenchmarkSpec
		if err := LoadYAMLTo(DefaultConfig, &spec); err != nil {
			return err
		}

		// Register
		regRes, err := client.RegisterGoBenchmark(context.Background(), &spec)
		if err != nil {
			return errors.Wrap(err, "failed to register")
		}
		log.Infof("registered job id %d spec id %d", regRes.JobId, regRes.SpecId)

		// Run with redirect
		if err := runGobench(spec.Command); err != nil {
			return errors.Wrap(err, "error run benchmark")
		}

		// Report
		result, err := gobench.ParseFile(spec.Report.Input)
		if err != nil {
			return errors.Wrap(err, "error parse benchmark output")
		}
		log.Infof("found %d results", len(result))

		// TODO: Add server side methods for report
		return nil
	},
}

func runGobench(command *bhpb.GoBenchmarkCommandSpec) error {
	cmd := exec.Command("sh", "-c", command.Command)
	var buf bytes.Buffer
	cmdout := io.MultiWriter(&buf, os.Stdout)
	cmd.Stdout = cmdout
	if err := cmd.Run(); err != nil {
		return err
	}
	if err := ioutil.WriteFile(command.Output, buf.Bytes(), 0664); err != nil {
		return err
	}
	return nil
}

func LoadYAMLTo(f string, v interface{}) error {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	if err := yaml.UnmarshalStrict(b, v); err != nil {
		return err
	}
	return nil
}
