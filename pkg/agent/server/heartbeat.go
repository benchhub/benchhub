package server

import (
	"context"
	"time"

	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"

	"github.com/benchhub/benchhub/pkg/agent/config"
	cpb "github.com/benchhub/benchhub/pkg/central/centralpb"
	"github.com/benchhub/benchhub/pkg/central/transport/grpc"
	pbc "github.com/benchhub/benchhub/pkg/common/commonpb"
)

const (
	// TODO: they should be configurable in beater struct
	registerTimeout  = 5 * time.Second
	heartbeatTimeout = 200 * time.Millisecond
)

// Beater keep posting the server about agent state and retrieve job status
type Beater struct {
	client     grpc.BenchHubCentralClient
	interval   time.Duration
	registered bool
	id         string

	registry     *Registry
	globalConfig config.ServerConfig

	log *dlog.Logger
}

func NewBeater(client grpc.BenchHubCentralClient, r *Registry) (*Beater, error) {
	interval, err := time.ParseDuration(r.Config.Heartbeat.Interval)
	if err != nil || interval <= 0 {
		return nil, errors.Errorf("invalid heartbeat interval config %d", interval)
	}
	b := &Beater{
		client:       client,
		interval:     interval,
		registry:     r,
		globalConfig: r.Config,
	}
	dlog.NewStructLogger(log, b)
	return b, nil
}

func (b *Beater) RunWithContext(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			// TODO: should we return nil or return context error?
			b.log.Infof("beater stop due to context finished, its error is %v", ctx.Err())
			return nil
		default:
			if !b.registered {
				// TODO: already exists error should be handled properly
				err := b.Register()
				if err != nil {
					b.log.Warnf("register failed %s, retry in %s", err.Error(), b.interval)
				} else {
					b.log.Infof("register success")
					b.registered = true
					// TODO: publish event inside process? the state machine now is just a thread safe struct
					b.registry.State.RegisterSuccess()
				}
			} else {
				// TODO: log every 100 requests (sampled logger), this should be supported by logging library,
				// keep a counter for it, zap seems to have it
				err := b.Beat()
				if err != nil {
					b.log.Warnf("heart beat failed %s", err.Error())
				}
				// TODO: need to switch to register
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
	b.id = res.Id
	b.log.Infof("register res id is %s", res.Id)
	return nil
}

func (b *Beater) Beat() error {
	c := b.client
	ctx, cancel := context.WithTimeout(context.Background(), heartbeatTimeout)
	defer cancel()
	req := &cpb.AgentHeartbeatReq{
		Id: b.id,
		Status: pbc.NodeStatus{
			State: b.registry.State.Current(),
		},
	}
	_, err := c.AgentHeartbeat(ctx, req)
	if err != nil {
		return errors.Wrap(err, "heart beat failed")
	}
	return nil
}
