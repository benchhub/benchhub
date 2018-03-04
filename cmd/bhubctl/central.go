package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	cpb "github.com/benchhub/benchhub/pkg/central/centralpb"
	mygrpc "github.com/benchhub/benchhub/pkg/central/transport/grpc"
	pbc "github.com/benchhub/benchhub/pkg/common/commonpb"
	"io/ioutil"
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

func (c *CentralCommand) SubmitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "submit",
		Short: "submit job",
		Long:  "Submit job for BenchHub to run",
		Run: func(cmd *cobra.Command, args []string) {
			c.mustCreateClient()
			if len(args) < 1 {
				log.Fatal("didn't provide spec file")
			}
			b, err := ioutil.ReadFile(args[0])
			if err != nil {
				log.Fatalf("failed to read file %s %v", args[0], err)
			}
			if res, err := c.client.SubmitJob(context.Background(), &cpb.SubmitJobReq{
				Spec: string(b),
			}); err != nil {
				log.Fatalf("submit job failed %v", err)
			} else {
				log.Infof("job submitted id is %s", res.Id)
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
	centralCmd.AddCommand(central.PingCmd(), central.SubmitCmd())
}
