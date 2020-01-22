package server

import (
	"bytes"
	"context"
	"time"

	dconfig "github.com/dyweb/gommon/config"
	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
	"github.com/benchhub/benchhub/pkg/central/job"
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
	var spec pb.JobSpec
	if err := dconfig.LoadYAMLDirectFromStrict(bytes.NewReader([]byte(req.Spec)), &spec); err != nil {
		errMsg := errors.Wrap(err, "can't parse YAML spec spec").Error()
		return &pb.SubmitJobRes{
			Error: &pb.Error{
				Code:    pb.ErrorCode_INVALID_CONFIG,
				Message: errMsg,
			},
		}, nil
	}
	// TODO: validate spec spec
	if id, err := srv.meta.AddJobSpec(spec); err != nil {
		errMsg := errors.Wrap(err, "can't add spec spec to meta store").Error()
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
	for {
		select {
		case <-ctx.Done():
			// TODO: tell all managers to stop?
			j.log.Infof("job poller stop due to context finished, its error is %v", ctx.Err())
			return nil
		default:
			j.processPending()
			time.Sleep(j.interval)
		}
	}
	j.log.Infof("stop job controller %s duration %s", time.Now(), time.Now().Sub(start))
	return nil
}

func (j *JobPoller) processPending() {
	meta := j.registry.Meta
	scheduler := j.registry.Scheduler

	spec, empty, err := meta.GetPendingJobSpec()
	if err != nil {
		j.log.Warnf("failed to get pending job %v", err)
		return
	} else if empty {
		return
	}
	j.log.Infof("start processing job %s", spec.Id)
	nodes, err := meta.ListNodes()
	if err != nil {
		j.log.Warnf("failed to list nodes %v", err)
		meta.PushbackJobSpec(spec.Id, spec)
		return
	}
	j.log.Infof("total %d nodes", len(nodes))
	assigned, err := scheduler.AssignNode(nodes, spec.NodeAssignments)
	if err != nil {
		j.log.Warnf("failed to assign nodes %v", err)
		meta.PushbackJobSpec(spec.Id, spec)
		return
	}
	j.log.Infof("assign finished")
	mgr := job.NewManager()
	mgr.SetSpec(spec)
	mgr.SetAssignedNodes(assigned)
	if err := j.registry.AddJob(mgr); err != nil {
		j.log.Warnf("failed to add job to registry %v", err)
		meta.PushbackJobSpec(spec.Id, spec)
		return
	}
	// TODO: need to start the job manager in background
	j.log.Infof("created manager for job %s", spec.Id)
}
