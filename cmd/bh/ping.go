package main

import (
	"context"

	"github.com/benchhub/benchhub/bhpb"
	"github.com/benchhub/benchhub/pkg/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// ping.go ping the server

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "ping server",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: use arg
		// TODO: cert
		conn, err := grpc.Dial(server.DefaultAddr, grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()
		client := bhpb.NewBenchHubClient(conn)
		res, err := client.Ping(context.Background(), &bhpb.PingRequest{Content: "hi"})
		if err != nil {
			return err
		}
		log.Infof("Ping response %s", res.Content)
		return nil
	},
}
