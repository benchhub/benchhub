package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/benchhub/benchhub/lib/waitforit"
)

var flags = flag.NewFlagSet("waitforit", flag.ExitOnError)
var showHelp = flags.Bool("h", false, "display help")
var timeout = flags.Duration("t", 10*time.Second, "time out for each connect")
var totalTimeout = flags.Duration("tt", 60*time.Second, "total time out")
var retry = flags.Int("r", 5, "number of retries for each host, negative for infinite retry")

type hostFlags []string

var hosts hostFlags

func (i *hostFlags) String() string {
	return fmt.Sprint(*i)
}

func (i *hostFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func help() {
	flags.Usage()
	// TODO: can we wait for UDP ?....
	example := `
Example:
  - Wait for two server using default config, return 0 when both of them can be connected
    waitforit -w tcp://localhost:9042 -w http://localhost:8080
`
	os.Stderr.Write([]byte(example))
	os.Exit(1)
}

func parseFlags(args []string) {
	if err := flags.Parse(args); err != nil {
		log.Fatal(err)
	}
	if *showHelp {
		help()
	}
}

func main() {
	if len(os.Args) < 2 {
		help()
	}
	parseFlags(os.Args[1:])
	var urls []url.URL
	for _, host := range hosts {
		if u, err := url.Parse(host); err != nil {
			log.Fatalf("invalid host %v", err)
		} else {
			urls = append(urls, *u)
		}
	}
	if d, err := waitforit.WaitSockets(context.Background(), urls, *totalTimeout, *timeout, *retry); err != nil {
		log.Fatalf("error %v", err)
	} else {
		log.Printf("finished after %s\n", d)
	}
}

func init() {
	flags.Var(&hosts, "host", "hosts to wait")
	flags.Var(&hosts, "w", "(alias of host) hosts to wait")
}
