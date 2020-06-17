package main

import (
	"context"
	"os"
	"sync"

	"github.com/benchhub/benchhub/exp/qaq16/config"
	"github.com/benchhub/benchhub/exp/qaq16/docker"
	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"
)

var (
	logReg = dlog.NewRegistry()
	log    = logReg.Logger()
)

func main() {
	args := os.Args
	// TODO: use gommon/dcli, don't want to use cobra anymore
	if len(args) < 2 {
		log.Fatal("must provide context e.g. qaq16 go")
		return
	}
	ctx := args[1]
	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

// TODO: allow dry run, print config and exit
func run(contextName string) error {
	cfg, err := config.Read("qaq15.yml")
	if err != nil {
		return err
	}
	// select context
	var runContext config.Context
	for _, ctx := range cfg.Contexts {
		if ctx.Name == contextName {
			runContext = ctx
			break
		}
	}
	if runContext.Name == "" {
		return errors.Errorf("context not found no %s", contextName)
	}

	// TODO: replace parameters in envs
	// TODO: expands container extend

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	merr := errors.NewMultiErrSafe()
	var wg sync.WaitGroup
	wg.Add(len(cfg.Containers) + 1) // TODO: use count of actual runnable container
	// Score
	go func() {
		defer wg.Done()

		log.Info("TODO: run score in host")
	}()
	// Containers
	for _, container := range cfg.Containers {
		container := container
		container.Image = runContext.Image
		go func() {
			defer wg.Done()

			if err := docker.Run(ctx, container); err != nil {
				merr.Append(err)
				log.Error(err)
				cancel() // TODO: is cancel go routine safe?
			}
		}()
	}
	wg.Wait()
	return merr.ErrorOrNil()
}
