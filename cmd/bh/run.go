package main

import (
	"context"

	"github.com/benchhub/benchhub/pkg/gobench"
	"github.com/spf13/cobra"
)

// run.go allocate a job id, run the benchmark and report result

const DefaultConfig = "benchhub.yml"

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run benchmark",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := mustDefaultClient()
		ctx := context.Background()
		runner, err := gobench.NewRunner(client, DefaultConfig)
		if err != nil {
			return err
		}
		if err := runner.Register(ctx); err != nil {
			return err
		}
		if err := runner.Run(ctx); err != nil {
			return err
		}
		if err := runner.Report(ctx); err != nil {
			return err
		}
		return nil
	},
}
