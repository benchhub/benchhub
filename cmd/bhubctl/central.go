package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	mygrpc "github.com/benchhub/benchhub/pkg/central/transport/grpc"
	pbc "github.com/benchhub/benchhub/pkg/common/commonpb"
)

const (
	localCentralAddr = "localhost:6081"
)

var central *CentralCommand

type CentralCommand struct {
	addr   string
	client mygrpc.BenchHubCentralClient
}

var centralCmd = &cobra.Command{
	Use:     "central",
	Aliases: []string{"c"},
	Short:   "benchub central",
	Long:    "Communicate with BenchHub central",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

func (c *CentralCommand) PingCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "ping",
		Short: "ping central",
		Long:  "Ping BenchHub central using gRPC",
		Run: func(cmd *cobra.Command, args []string) {
			c.mustCreateClient()
			host, _ := os.Hostname()
			if res, err := c.client.Ping(context.Background(), &pbc.Ping{Message: "ping from " + host}); err != nil {
				log.Fatal(err)
			} else {
				log.Infof("ping finished central response is %s", res.Message)
			}
		},
	}
}

func (c *CentralCommand) mustCreateClient() {
	if c.client != nil {
		return
	}
	if conn, err := grpc.Dial(c.addr, grpc.WithInsecure()); err != nil {
		log.Fatalf("can't dial %v", err)
	} else {
		c.client = mygrpc.NewClient(conn)
	}
}

func init() {
	central := &CentralCommand{
		addr: localCentralAddr,
	}
	centralCmd.AddCommand(central.PingCmd())
}
