package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/benchhub/benchhub/bhpb"
	"github.com/benchhub/benchhub/pkg/server"
	icli "github.com/dyweb/go.ice/cli"
	dlog "github.com/dyweb/gommon/log"
	"google.golang.org/grpc"
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
	root.AddCommand(runCmd)
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func mustDefaultClient() bhpb.BenchHubClient {
	// TODO: cert
	// TODO: close the connection? what happens if we don't?
	conn, err := grpc.Dial(server.DefaultAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to create default client %s", err)
		return nil
	}
	return bhpb.NewBenchHubClient(conn)
}
