package main

import (
	"fmt"
	"os"
	"runtime"

	icli "github.com/at15/go.ice/ice/cli"
	"github.com/benchhub/benchhub/lib/monitor/util/logutil"
)

const (
	myname = "bhmon"
)

// TODO: we should not use lib log as top registry, but they are same struct, so it should work
var log = logutil.Registry

var (
	version   string
	commit    string
	buildTime string
	buildUser string
	goVersion = runtime.Version()
)

var buildInfo = icli.BuildInfo{Version: version, Commit: commit, BuildTime: buildTime, BuildUser: buildUser, GoVersion: goVersion}

func main() {
	cli := icli.New(
		icli.Name(myname),
		icli.Description("BenchHub monitor lib test command"),
		icli.Version(buildInfo),
		icli.LogRegistry(log),
	)
	root := cli.Command()
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
