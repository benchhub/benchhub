package main

import (
	"context"
	"fmt"

	"github.com/benchhub/benchhub/bhpb"
	"github.com/benchhub/benchhub/core/config/defaults"
	"github.com/dyweb/gommon/dcli"
	"google.golang.org/grpc"
)

func CmdUser() *dcli.Cmd {
	return &dcli.Cmd{
		Name: "user",
		Run: func(ctx context.Context) error {
			log.Info("TODO: user")
			addr := fmt.Sprintf("localhost:%d", defaults.ServerPort)
			// TODO: maybe use tls even for localhost testing?
			conn, err := grpc.Dial(addr, grpc.WithInsecure())
			if err != nil {
				return err
			}
			client := bhpb.NewUserServiceClient(conn)
			user, err := client.GetUser(ctx, bhpb.NewName("at15"))
			if err != nil {
				return err
			}
			log.Infof("%d %s %s %s", user.Id, user.Name, user.FullName, user.Email)
			return nil
		},
	}
}
