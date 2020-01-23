package main

import (
	"context"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/benchhub/benchhub/bhpb"
	"github.com/dyweb/gommon/errors"
	"github.com/spf13/cobra"
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
		regRes, err := client.RegisterGoBenchmark(context.Background(), &spec)
		if err != nil {
			return errors.Wrap(err, "failed to register")
		}
		log.Infof("registered job id %d spec id %d", regRes.JobId, regRes.SpecId)
		return nil
	},
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
