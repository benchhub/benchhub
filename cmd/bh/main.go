package main

import (
	"context"

	"github.com/dyweb/gommon/dcli"
	dlog "github.com/dyweb/gommon/log"
)

var logReg = dlog.NewRegistry()
var log = logReg.NewLogger()

func main() {
	root := &dcli.Cmd{
		Name: "bh",
		Run: func(ctx context.Context) error {
			log.Info("bh does nothing")
			return nil
		},
		Children: []dcli.Command{
			CmdCore(),
			// Actions
			CmdRegister(),
			// Services
			CmdUser(),
		},
	}
	dcli.RunApplication(root)
}
