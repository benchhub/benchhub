package main

import (
	"context"
	"fmt"

	"github.com/benchhub/benchhub/core/config/defaults"
	"github.com/benchhub/benchhub/core/server"
	"github.com/dyweb/gommon/dcli"
)

func CmdCore() *dcli.Cmd {
	return &dcli.Cmd{
		Name: "core",
		Run: func(ctx context.Context) error {
			return dcli.NewErrHelpOnly("core")
		},
		Children: []dcli.Command{
			&dcli.Cmd{
				Name: "server",
				Run: func(ctx context.Context) error {
					// TODO: start and stop server in sub command
					addr := fmt.Sprintf("localhost:%d", defaults.ServerPort)
					srv, err := server.New(addr)
					if err != nil {
						return err
					}
					return srv.Run(ctx)
				},
			},
		},
	}
}
