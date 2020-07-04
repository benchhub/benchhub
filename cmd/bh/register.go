package main

import (
	"context"

	"github.com/dyweb/gommon/dcli"
)

// register.go defines command for bh register

func CmdRegister() *dcli.Cmd {
	return &dcli.Cmd{
		Name: "register",
		Run: func(ctx context.Context) error {
			log.Info("TODO: register")
			return nil
		},
	}
}
