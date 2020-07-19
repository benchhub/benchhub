package main

import (
	"context"

	"github.com/dyweb/gommon/util/fsutil"

	"github.com/benchhub/benchhub/lib/tqbuilder/generator"
	"github.com/dyweb/gommon/dcli"
	dlog "github.com/dyweb/gommon/log"
)

var (
	logReg = dlog.NewRegistry()
	log    = logReg.NewLogger()
)

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
					return genSchema()
				},
			},
		},
	}
	dcli.RunApplication(root)
}

const bhRepo = "github.com/benchhub/benchhub"

func genSchema() error {
	res, err := generator.Walk("core/services")
	if err != nil {
		return err
	}
	const dstDir = "build/generated/tqbuilder/ddl"
	ddlMain, err := fsutil.CreateFileAndPath(dstDir, "main.go")
	if err != nil {
		return err
	}
	defer ddlMain.Close()
	if err := generator.GenDDLMain(ddlMain, bhRepo, res.DDLs); err != nil {
		return err
	}
	log.Infof("TODO: DMLS %v", res.DMLS)
	return nil
}
