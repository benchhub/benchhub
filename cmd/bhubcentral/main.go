package main

import (
	"fmt"
	"os"
	"runtime"

	icli "github.com/at15/go.ice/ice/cli"

	"github.com/benchhub/benchhub/pkg/central/config"
	"github.com/benchhub/benchhub/pkg/util/logutil"
)

const (
	myname = "bhubcentral"
)

var log = logutil.Registry

var (
	version   string
	commit    string
	buildTime string
	buildUser string
	goVersion = runtime.Version()
)

var buildInfo = icli.BuildInfo{Version: version, Commit: commit, BuildTime: buildTime, BuildUser: buildUser, GoVersion: goVersion}
var cli *icli.Root
var cfg config.ServerConfig

func main() {
	cli = icli.New(
		icli.Name(myname),
		icli.Description("BenchHub central"),
		icli.Version(buildInfo),
		icli.LogRegistry(log),
	)
	root := cli.Command()
	root.AddCommand(serveCmd)
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func mustLoadConfig() {
	if err := cli.LoadConfigTo(&cfg); err != nil {
		log.Fatal(err)
	}
}
