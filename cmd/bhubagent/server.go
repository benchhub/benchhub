package main

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	igrpc "github.com/at15/go.ice/ice/transport/grpc"
	"github.com/benchhub/benchhub/pkg/agent/server"
	mygrpc "github.com/benchhub/benchhub/pkg/agent/transport/grpc"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start agent daemon",
	Long:  "Start BenchHub agent daemon with gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		mustLoadConfig()
		srv, err := server.NewGrpcServer()
		if err != nil {
			log.Fatal(err)
		}
		grpcSrv, err := igrpc.NewServer(cfg.Grpc, func(s *grpc.Server) {
			mygrpc.RegisterBenchHubAgentServer(s, srv)
		})
		if err != nil {
			log.Fatal(err)
		}
		if err := grpcSrv.Run(); err != nil {
			log.Fatal(err)
		}
	},
}
