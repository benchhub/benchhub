package main

import (
	"context"

	"github.com/benchhub/benchhub/bhpb"
	"github.com/spf13/cobra"
)

// ping.go ping the server

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "ping server",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: use arg to pick remote
		client := mustDefaultClient()
		res, err := client.Ping(context.Background(), &bhpb.PingRequest{Content: "hi"})
		if err != nil {
			return err
		}
		log.Infof("Ping response %s", res.Content)
		return nil
	},
}
