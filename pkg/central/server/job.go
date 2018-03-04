package server

import (
	"context"

	dlog "github.com/dyweb/gommon/log"
	"time"
)

type JobController struct {
	registry *Registry
	log      *dlog.Logger
}

func NewJobController() (*JobController, error) {
	j := &JobController{}
	dlog.NewStructLogger(log, j)
	return j, nil
}

func (j *JobController) RunWithContext(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			// TODO: should we return nil or return context error?
			// TODO: we should tell all the agent to abort job since central is shut down?
			j.log.Infof("job controller stop due to context finished, its error is %v", ctx.Err())
			return nil
		default:
			// TODO: pull for store if there is any pending job
			// TODO: poll duration should be configurable
			time.Sleep(1 * time.Second)
		}
	}
}
