package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	mygrpc "github.com/benchhub/benchhub/pkg/agent/transport/grpc"
	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

const (
	localAgentAddr = "localhost:6082"
)

var agent *AgentCommand

type AgentCommand struct {
	addr   string
	client mygrpc.BenchHubAgentClient
}

func (c *AgentCommand) PingCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "ping",
		Short: "ping agent",
		Long:  "Ping BenchHub agent using gRPC",
		Run: func(cmd *cobra.Command, args []string) {
			c.mustCreateClient()
			host, _ := os.Hostname()
			if res, err := c.client.Ping(context.Background(), &pb.Ping{Message: "ping from " + host}); err != nil {
				log.Fatal(err)
			} else {
				log.Infof("ping finished agent response is %s", res.Message)
			}
		},
	}
}

func (c *AgentCommand) mustCreateClient() {
	log.Infof("host is %s", c.addr)
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
	agent = &AgentCommand{
		addr: localAgentAddr,
	}
}
