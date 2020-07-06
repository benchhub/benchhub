package main

import (
	"context"
	"fmt"

	"github.com/benchhub/benchhub/core/config/defaults"
	"github.com/benchhub/benchhub/core/server"
	"github.com/dyweb/gommon/dcli"
)

func CmdServer() *dcli.Cmd {
	// TODO: start and stop server in sub command
	return &dcli.Cmd{
		Name: "server",
		Run: func(ctx context.Context) error {
			log.Info("TODO: server")
			addr := fmt.Sprintf("localhost:%d", defaults.ServerPort)
			srv, err := server.New(addr)
			if err != nil {
				return err
			}
			return srv.Run(ctx)
		},
	}
}
