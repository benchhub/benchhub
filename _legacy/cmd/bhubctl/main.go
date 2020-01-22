package main

import (
	"fmt"
	"os"
	"runtime"

	icli "github.com/at15/go.ice/ice/cli"
	"github.com/benchhub/benchhub/pkg/util/logutil"
)

const (
	myname = "bhubctl"
)

// TODO: should put logic in the manager struct in pkg/ctl instead of scatter in cmd
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
		icli.Description("BenchHub central client cli"),
		icli.Version(buildInfo),
		icli.LogRegistry(log),
	)
	root := cli.Command()
	root.PersistentFlags().StringVar(&central.addr, "host", localCentralAddr, "Host of central server")
	root.AddCommand(central.PingCmd(), central.SubmitCmd())
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
