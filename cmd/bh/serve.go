package main

import (
	"context"

	"github.com/benchhub/benchhub/pkg/server"
	"github.com/spf13/cobra"
)

// serve.go starts a benchhub server

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start server",
	RunE: func(cmd *cobra.Command, args []string) error {
		var cfg server.Config
		if err := cli.LoadConfigTo(&cfg); err != nil {
			return err
		}
		srv, err := server.New(cfg)
		if err != nil {
			return err
		}
		return srv.Run(context.Background())
	},
}
