package main

import (
	"context"
	"os"
	"path/filepath"

	"github.com/benchhub/benchhub/lib/tqbuilder/generator"
	"github.com/dyweb/gommon/dcli"
	"github.com/dyweb/gommon/errors"
	"github.com/dyweb/gommon/util/fsutil"
)

// schema.go is the command that generates RDBMS binding using tqbuilder

func schemaCmd() dcli.Command {
	return &dcli.Cmd{
		Name: "schema",
		Run: func(ctx context.Context) error {
			return &dcli.ErrHelpOnlyCommand{}
		},
		Children: []dcli.Command{
			&dcli.Cmd{
				Name: "generate",
				Run: func(ctx context.Context) error {
					return genSchema()
				},
			},
			&dcli.Cmd{
				Name: "clean",
				Run: func(ctx context.Context) error {
					return cleanSchema()
				},
			},
		},
	}
}

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
