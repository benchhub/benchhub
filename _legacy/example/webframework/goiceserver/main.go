package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	icli "github.com/at15/go.ice/ice/cli"
	iconfig "github.com/at15/go.ice/ice/config"
	ihttp "github.com/at15/go.ice/ice/transport/http"
	goicelog "github.com/at15/go.ice/ice/util/logutil"
	dlog "github.com/dyweb/gommon/log"
	"github.com/spf13/cobra"
)

const (
	myname = "goiceserver"
)

var log = dlog.NewApplicationLogger()

var (
	version   string
	commit    string
	buildTime string
	buildUser string
	goVersion = runtime.Version()
)

var buildInfo = icli.BuildInfo{Version: version, Commit: commit, BuildTime: buildTime, BuildUser: buildUser, GoVersion: goVersion}
var cli *icli.Root
var cfg *Config

type Config struct {
	Http iconfig.HttpServerConfig `yaml:"http"`
}

type HttpServer struct {
}

func (s *HttpServer) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func (s *HttpServer) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", s.Ping)
	return mux
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server",
	Run: func(cmd *cobra.Command, args []string) {
		mustLoadConfig()
		srv := &HttpServer{}
		transport, err := ihttp.NewServer(cfg.Http, srv.Handler(), nil)
		if err != nil {
			log.Fatal(err)
		}
		if err := transport.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func main() {
	cli = icli.New(
		icli.Name(myname),
		icli.Description("Test server for framework performance"),
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

func init() {
	log.AddChild(goicelog.Registry)
}
