package main

import (
	"context"
	"github.com/benchhub/benchhub/lib/tqbuilder/generator"

	"github.com/dyweb/gommon/dcli"
	dlog "github.com/dyweb/gommon/log"
)

var logReg = dlog.NewRegistry()
var log = logReg.NewLogger()

func main() {
	root := &dcli.Cmd{
		Name: "bhgen",
		Run: func(ctx context.Context) error {
			log.Info("bhgen should print help, but gommon/dcli does not support it yet")
			return nil
		},
		Children: []dcli.Command{
			&dcli.Cmd{
				Name: "schema",
				Run: func(ctx context.Context) error {
					return schema()
				},
			},
		},
	}
	dcli.RunApplication(root)
}

func schema() error {
	res, err := generator.Walk("core/services")
	if err != nil {
		return err
	}
	log.Infof("TODO: %v", res.DDLs)
	return nil
}
