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

const minPollInterval = time.Second / 2

// JobPoller get job for store and create job managers to run them
type JobPoller struct {
	registry *Registry
	interval time.Duration

	log *dlog.Logger
}

func NewJobPoller(r *Registry, pollInterval time.Duration) (*JobPoller, error) {
	if pollInterval < minPollInterval {
		pollInterval = minPollInterval
	}
	j := &JobPoller{
		registry: r,
		interval: pollInterval,
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
	j.log.Infof("start job controller %s", time.Now())
	start := time.Now()
	meta := j.registry.Meta
	interval := j.interval
	for {
		select {
		case <-ctx.Done():
			// TODO: tell all managers to stop?
			j.log.Infof("job poller stop due to context finished, its error is %v", ctx.Err())
			return nil
		default:
			job, empty, err := meta.GetPendingJob()
			if err != nil {
				log.Warnf("failed to get pending job %v", err)
				goto SLEEP
			}
			if empty {
				goto SLEEP
			}
			log.Infof("TODO: deal with job %s", job.Id)
		SLEEP:
			time.Sleep(interval)
		}
	}
	j.log.Infof("stop job controller %s duration %s", time.Now(), time.Now().Sub(start))
	return nil
}
