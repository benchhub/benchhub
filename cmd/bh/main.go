package main

import (
	"fmt"
	"os"
	"runtime"

	icli "github.com/dyweb/go.ice/cli"
	dlog "github.com/dyweb/gommon/log"
)

const (
	myname = "bh"
)

var logReg = dlog.NewRegistry()
var log = logReg.Logger()

var (
	version   string
	commit    string
	buildTime string
	buildUser string
	goVersion = runtime.Version()
)

var buildInfo = icli.BuildInfo{Version: version, Commit: commit, BuildTime: buildTime, BuildUser: buildUser, GoVersion: goVersion}

var cli *icli.Root

func main() {
	cli = icli.New(
		icli.Name(myname),
		icli.Description("BenchHub"),
		icli.Version(buildInfo),
	)
	root := cli.Command()
	root.AddCommand(serveCmd)
	root.AddCommand(pingCmd)
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
