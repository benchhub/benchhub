package server

import (
	"bytes"
	"context"
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

// SubmitJob handles job spec submitted by client and return an id to it
func (srv *GrpcServer) SubmitJob(ctx context.Context, req *pb.SubmitJobReq) (*pb.SubmitJobRes, error) {
	var job pb.JobSpec
	if err := dconfig.LoadYAMLDirectFromStrict(bytes.NewReader([]byte(req.Spec)), &job); err != nil {
		errMsg := errors.Wrap(err, "can't parse YAML job spec").Error()
		return &pb.SubmitJobRes{
			Error: &pb.Error{
				Code:    pb.ErrorCode_INVALID_CONFIG,
				Message: errMsg,
			},
		}, nil
	}
	// TODO: validate job spec
	if id, err := srv.meta.AddJobSpec(job); err != nil {
		errMsg := errors.Wrap(err, "can't add job spec to meta store").Error()
		return &pb.SubmitJobRes{
			Error: &pb.Error{
				Code:    pb.ErrorCode_STORE_ERROR,
				Message: errMsg,
			},
		}, nil
	} else {
		return &pb.SubmitJobRes{Id: id}, nil
	}
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
