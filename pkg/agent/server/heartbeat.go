package server

import (
	"context"
	"time"

	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"

	cpb "github.com/benchhub/benchhub/pkg/central/centralpb"
	"github.com/benchhub/benchhub/pkg/central/transport/grpc"
	"github.com/benchhub/benchhub/pkg/common/nodeutil"
)

const (
	// TODO: they should be configurable in beater struct
	registerTimeout  = 5 * time.Second
	heartbeatTimeout = 20 * time.Millisecond
)

// Beater keep posting the server about agent state and retrieve job status
type Beater struct {
	client grpc.BenchHubCentralClient
	log    *dlog.Logger
}

func NewBeater(client grpc.BenchHubCentralClient) *Beater {
	b := &Beater{client: client}
	dlog.NewStructLogger(log, b)
	return b
}

func (b *Beater) RunWithContext(ctx context.Context) error {
	// TODO: based on agent state, either register or heartbeat
	b.log.Warn(b.Register())
	return nil
}

func (b *Beater) Register() error {
	c := b.client
	ctx, cancel := context.WithTimeout(context.Background(), registerTimeout)
	defer cancel()
	node, err := nodeutil.GetNode()
	// TODO: update bindAddr, ip, port, etc.
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
