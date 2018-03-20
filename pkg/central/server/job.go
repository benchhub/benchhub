package server

import (
	"bytes"
	"context"
	"fmt"
	"sync/atomic"
	"time"

	dconfig "github.com/dyweb/gommon/config"
	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

// JobPoller get job for store and create job managers to run them
type JobPoller struct {
	registry *Registry
	log      *dlog.Logger
}

func NewJobPoller(r *Registry) (*JobPoller, error) {
	j := &JobPoller{
		registry: r,
	}
	dlog.NewStructLogger(log, j)
	return j, nil
}

func (srv *GrpcServer) SubmitJob(ctx context.Context, req *pb.SubmitJobReq) (*pb.SubmitJobRes, error) {
	var job pb.JobSpec
	if err := dconfig.LoadYAMLDirectFromStrict(bytes.NewReader([]byte(req.Spec)), &job); err != nil {
		return nil, errors.Wrap(err, "can't parse YAML job spec")
	}
	// TODO: implement the validate logic
	//if err := job.Validate(); err != nil {
	//	return nil, errors.Wrap(err, "invalid job spec")
	//}
	// TODO: wrap this in store, store should return an id for job ...
	// FIXME: we are just using project name + a global counter ...
	atomic.AddInt64(&srv.c, 1)
	id := fmt.Sprintf("%s-%d", job.Name, atomic.LoadInt64(&srv.c))
	srv.log.Infof("got job %s id %s", job.Name, id)
	if err := srv.meta.AddJobSpec(id, job); err != nil {
		return nil, errors.Wrap(err, "can't add job spec to store")
	}
	return &pb.SubmitJobRes{Id: id}, nil
}

func (j *JobPoller) RunWithContext(ctx context.Context) error {
	j.log.Info("start job controller")
	meta := j.registry.Meta
	for {
		select {
		case <-ctx.Done():
			// TODO: tell all managers to stop?
			j.log.Infof("job poller stop due to context finished, its error is %v", ctx.Err())
			return nil
		default:
			job, empty, err := meta.GetPendingJob()
			if empty {
				// do nothing
			} else if err != nil {
				log.Warnf("failed to get pending job %v", err)
			} else {
				// TODO: get nodes and assign jobs
				// TODO: put back if failed to schedule
				log.Infof("TODO: deal with job %s", job.Id)
			}
			// TODO: poll duration should be configurable
			time.Sleep(1 * time.Second)
		}
	}
}
