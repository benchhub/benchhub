package main

import (
	"context"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/dyweb/gommon/log/handlers/cli"

	"github.com/benchhub/benchhub/exp/qaq16/config"
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

func mergeEnvs(base []config.Env, ext []config.Env) []config.Env {
	var merged []config.Env
	merged = append(merged, base...)
	merged = append(merged, ext...)
	// TODO: remove dup
	return merged
}

func extendContainer(containers []config.Container) ([]config.Container, error) {
	abstracts := make(map[string]config.Container)
	var unresolved []config.Container
	var resolved []config.Container
	for _, c := range containers {
		switch {
		case c.Abstract:
			abstracts[c.Name] = c
		case c.Extends == "":
			resolved = append(resolved, c)
		case c.Extends != "":
			unresolved = append(unresolved, c)
		default:
			// should never happen
			return nil, errors.Errorf("container must be abstract, extends or final %s", c.Name)
		}
	}
	for _, c := range unresolved {
		base, ok := abstracts[c.Extends]
		if !ok {
			return nil, errors.Errorf("container base not found %s wants %s", c.Name, c.Extends)
		}
		// simply copy everything
		c.Image = base.Image
		c.Resource = base.Resource
		c.Envs = mergeEnvs(base.Envs, c.Envs)
		resolved = append(resolved, c)
	}
	return resolved, nil
}

const paramPrefix = "param."

func resolveEnv(c *config.Container, params []config.Parameter) error {
	pmap := make(map[string]int)
	for _, p := range params {
		pmap[p.Name] = p.Default
	}
	for i, env := range c.Envs {
		if strings.HasPrefix(env.Value, paramPrefix) {
			key := strings.TrimPrefix(env.Value, paramPrefix)
			v, ok := pmap[key]
			if !ok {
				return errors.Errorf("param not found container %s requires %s", c.Name, key)
			}
			env.Value = strconv.Itoa(v)
			c.Envs[i] = env
		}
	}
	return nil
}

// TODO: allow dry run, print config and exit
func run(contextName string) error {
	cfg, err := config.Read("qaq15.yml")
	if err != nil {
		return err
	}
	// select context
	var selectedContext config.Context
	for _, ctx := range cfg.Contexts {
		if ctx.Name == contextName {
			selectedContext = ctx
			break
		}
	}
	if selectedContext.Name == "" {
		return errors.Errorf("context not found no %s", contextName)
	}

	containers, err := extendContainer(cfg.Containers)
	if err != nil {
		return err
	}
	for i := range containers {
		if err := resolveEnv(&containers[i], cfg.Parameters); err != nil {
			return err
		}
	}

	now := time.Now()
	logPrefix, err := NewLogDir(cfg.Data, now)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(len(containers) + 1)
	merr := errors.NewMultiErrSafe()

	// Run containers
	for _, container := range containers {
		container := container
		container.Image = selectedContext.Image
		go func() {
			defer wg.Done()

			execCtx := ExecContext{log: FormatLog(logPrefix, container.Name)}
			if err := RunContainer(ctx, container, execCtx); err != nil {
				merr.Append(err)
				log.Error(err)
				cancel() // TODO: is cancel go routine safe?
			}
		}()
	}
	// FIXME: hack to wait container ready
	time.Sleep(1 * time.Second)

	// Run score
	go func() {
		defer wg.Done()

		execCtx := ExecContext{log: FormatLog(logPrefix, "score")}
		if err := RunScore(ctx, cfg.Score, execCtx); err != nil {
			merr.Append(err)
			log.Error(err)
			cancel()
		}
	}()

	wg.Wait()
	return merr.ErrorOrNil()
}

func init() {
	dlog.SetHandler(cli.New(os.Stderr, false))
}
