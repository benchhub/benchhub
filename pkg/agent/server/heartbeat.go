package server

import (
	"context"
	"time"

	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"

	"github.com/benchhub/benchhub/pkg/agent/config"
	cpb "github.com/benchhub/benchhub/pkg/central/centralpb"
	"github.com/benchhub/benchhub/pkg/central/transport/grpc"
)

const (
	// TODO: they should be configurable in beater struct
	registerTimeout  = 5 * time.Second
	heartbeatTimeout = 20 * time.Millisecond
)

// Beater keep posting the server about agent state and retrieve job status
type Beater struct {
	client       grpc.BenchHubCentralClient
	interval     time.Duration
	globalConfig config.ServerConfig
	registered   bool
	log          *dlog.Logger
}

func NewBeater(client grpc.BenchHubCentralClient, interval time.Duration, cfg config.ServerConfig) *Beater {
	b := &Beater{
		client:       client,
		interval:     interval,
		globalConfig: cfg,
	}
	dlog.NewStructLogger(log, b)
	return b
}

func (b *Beater) RunWithContext(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			// TODO: should we return nil or return context error?
			b.log.Infof("beater stop due to context finished, its error is %v", ctx.Err())
			return nil
		default:
			// TODO: based on agent state, either register or heartbeat
			if !b.registered {
				err := b.Register()
				if err != nil {
					b.log.Warnf("register failed %s, retry in %s", err.Error(), b.interval)
				} else {
					b.log.Infof("register success")
					b.registered = true
					// TODO: publish event inside process? need a place to know node's state
				}
			} else {
				// TODO: real heart beat logic
				b.log.Infof("TODO: heartbeat")
			}
			time.Sleep(b.interval)
		}
	}
	return nil
}

func (b *Beater) Register() error {
	c := b.client
	ctx, cancel := context.WithTimeout(context.Background(), registerTimeout)
	defer cancel()
	node, err := Node(b.globalConfig)
	if err != nil {
		return err
	}
	req := &cpb.RegisterAgentReq{
		Node: *node,
	}
	res, err := c.RegisterAgent(ctx, req)
	if err != nil {
		return errors.Wrap(err, "register failed")
	}
	b.log.Infof("register res id is %s", res.Id)
	return nil
}
