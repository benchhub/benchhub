package main

import (
	"github.com/spf13/cobra"

	"github.com/benchhub/benchhub/pkg/central/server"
	"github.com/benchhub/benchhub/pkg/common/nodeutil"

	// empty imports to enable providers
	_ "github.com/benchhub/benchhub/pkg/central/store/meta/mem"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start central daemon",
	Long:  "Start BenchHub central daemon with gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("node id is %s", nodeutil.UID())
		mustLoadConfig()
		mgr, err := server.NewManager(cfg)
		if err != nil {
			log.Fatal(err)
		}
		if err := mgr.Run(); err != nil {
			log.Fatal(err)
		}
	},
}
