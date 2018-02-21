package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	pb "github.com/benchhub/benchhub/pkg/agent/agentpb"
	mygrpc "github.com/benchhub/benchhub/pkg/agent/transport/grpc"
)

const (
	localAgentAddr = "localhost:6082"
)

var agent *AgentCommand

var agentCmd = &cobra.Command{
	Use:     "agent",
	Aliases: []string{"a"},
	Short:   "benchub agent",
	Long:    "Communicate with BenchHub agent",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

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
			if res, err := c.client.Ping(context.Background(), &pb.Ping{Name: host}); err != nil {
				log.Fatal(err)
			} else {
				log.Infof("ping finished agent host name is %s", res.Name)
			}
		},
	}
}

func (c *AgentCommand) mustCreateClient() {
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
	agentCmd.AddCommand(agent.PingCmd())
}
