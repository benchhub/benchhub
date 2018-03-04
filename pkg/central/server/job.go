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

func NewJobController(r *Registry) (*JobController, error) {
	j := &JobController{
		registry: r,
	}
	dlog.NewStructLogger(log, j)
	return j, nil
}

func (j *JobController) RunWithContext(ctx context.Context) error {
	j.log.Info("start job controller")
	meta := j.registry.Meta
	for {
		select {
		case <-ctx.Done():
			// TODO: should we return nil or return context error?
			// TODO: we should tell all the agent to abort job since central is shut down?
			j.log.Infof("job controller stop due to context finished, its error is %v", ctx.Err())
			return nil
		default:
			// TODO: pull for store if there is any pending job
			job, empty, err := meta.GetPendingJob()
			if empty {
				// do nothing
			} else if err != nil {
				log.Warnf("failed to get pending job %v", err)
			} else {
				// TODO: spec should contains id ...
				log.Infof("TODO: process job %s", job.Name)
			}
			// TODO: poll duration should be configurable
			time.Sleep(1 * time.Second)
		}
	}
}
