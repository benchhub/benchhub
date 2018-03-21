package main

import (
	"github.com/spf13/cobra"

	"github.com/benchhub/benchhub/pkg/agent/server"
	"github.com/benchhub/benchhub/pkg/common/nodeutil"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start agent daemon",
	Long:  "Start BenchHub agent daemon with gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("node id is %s", nodeutil.UID())
		mustLoadConfig()
		log.Infof("node role is %s", cfg.Node.Role)
		cfg.Grpc.Addr = grpcAddr
		cfg.Http.Addr = httpAddr
		mgr, err := server.NewManager(cfg)
		if err != nil {
			log.Fatal(err)
		}
		if err := mgr.Run(); err != nil {
			log.Fatal(err)
		}
	},
}
