package main

import (
	"context"
	"os"
	"path/filepath"

	"github.com/dyweb/gommon/errors"
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
				Children: []dcli.Command{
					&dcli.Cmd{
						Name: "clean",
						Run: func(ctx context.Context) error {
							return cleanSchema()
						},
					},
				},
			},
		},
	}
	dcli.RunApplication(root)
}

const bhRepo = "github.com/benchhub/benchhub"

// TODO: remove hard coded path and move the logic into tqbuilder/generator
func genSchema() error {
	res, err := generator.Walk("core/services", bhRepo)
	if err != nil {
		return err
	}
	const dstDir = "build/generated/tqbuilder/ddl"
	ddlMain, err := fsutil.CreateFileAndPath(dstDir, "main.go")
	if err != nil {
		return err
	}
	defer ddlMain.Close()
	if err := generator.GenDDLMain(ddlMain, res.DDLs); err != nil {
		return err
	}
	log.Infof("TODO: DMLS %v", res.DMLS)
	return nil
}

func cleanSchema() error {
	rmGoFile := func(path string, info os.FileInfo, err error) error {
		if fsutil.IsGoFile(info) {
			log.Debugf("Remove %s", path)
			return os.Remove(path)
		}
		return nil
	}
	merr := errors.NewMulti()
	merr.Append(filepath.Walk("build/generated/tqbuilder", rmGoFile))
	res, err := generator.Walk("core/services", bhRepo)
	if err != nil {
		return err
	}
	for _, ddl := range res.DDLs {
		merr.Append(filepath.Walk(ddl.OutputPath, rmGoFile))
	}
	return merr.ErrorOrNil()
}
