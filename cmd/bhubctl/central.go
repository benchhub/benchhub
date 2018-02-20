package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	pb "github.com/benchhub/benchhub/pkg/central/centralpb"
	mygrpc "github.com/benchhub/benchhub/pkg/central/transport/grpc"
)

var centralAddr = "localhost:6081"
var centralClient mygrpc.BenchHubCentralClient

var centralCmd = &cobra.Command{
	Use:   "central",
	Short: "benchub central",
	Long:  "Communicate with BenchHub central",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

var centralPingCmd = &cobra.Command{
	Use:   "ping",
	Short: "ping central",
	Long:  "Ping BenchHub central using gRPC",
	Run: func(cmd *cobra.Command, args []string) {
		mustCreateCentralClient()
		host, _ := os.Hostname()
		if res, err := centralClient.Ping(context.Background(), &pb.Ping{Name: host}); err != nil {
			log.Fatal(err)
		} else {
			log.Infof("ping finished central host name is %s", res.Name)
		}
	},
}

func mustCreateCentralClient() {
	if conn, err := grpc.Dial(centralAddr, grpc.WithInsecure()); err != nil {
		log.Fatalf("can't dial %v", err)
	} else {
		centralClient = mygrpc.NewClient(conn)
	}
}

func init() {
	centralCmd.AddCommand(centralPingCmd)
}
