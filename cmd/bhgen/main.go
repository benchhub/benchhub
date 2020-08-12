package main

import (
	"context"

	"github.com/dyweb/gommon/dcli"
	dlog "github.com/dyweb/gommon/log"
)

var (
	logReg = dlog.NewRegistry()
	log    = logReg.NewLogger()
)

// TODO: remove the hard coded import path for benchhub?
const bhRepo = "github.com/benchhub/benchhub"

func main() {
	root := &dcli.Cmd{
		Name: "bhgen",
		Run: func(ctx context.Context) error {
			return &dcli.ErrHelpOnlyCommand{}
		},
		Children: []dcli.Command{
			schemaCmd(),
		},
	}
	dcli.RunApplication(root)
}
