package gobench

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/benchhub/benchhub/bhpb"
	"github.com/dyweb/gommon/errors"
)

type Runner struct {
	spec   bhpb.GoBenchmarkSpec
	client bhpb.BenchHubClient

	// responses from server
	jobId  int64
	specId int64

	// record when running
	startTime time.Time
	endTime   time.Time
}

func NewRunner(bhclient bhpb.BenchHubClient, specFile string) (*Runner, error) {
	var spec bhpb.GoBenchmarkSpec
	// TODO: should validate input
	if err := LoadYAMLTo(specFile, &spec); err != nil {
		return nil, err
	}
	return &Runner{
		spec:   spec,
		client: bhclient,
	}, nil
}

// Register upload spec to benchhub and allocate ids
func (r *Runner) Register(ctx context.Context) error {
	regRes, err := r.client.GoBenchmarkRegisterJob(ctx, &r.spec)
	if err != nil {
		return err
	}

	r.jobId, r.specId = regRes.JobId, regRes.SpecId
	return nil
}

func (r *Runner) Run(ctx context.Context) error {
	if r.jobId == 0 {
		return errors.New("job is not registered, jobId is 0")
	}

	r.startTime = time.Now()
	command := r.spec.Command
	cmd := exec.Command("sh", "-c", command.Command)
	var buf bytes.Buffer
	cmdout := io.MultiWriter(&buf, os.Stdout)
	cmd.Stdout = cmdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "error running command %s", command.Command)
	}
	if err := ioutil.WriteFile(command.Output, buf.Bytes(), 0664); err != nil {
		return err
	}
	r.endTime = time.Now()
	return nil
}

func (r *Runner) Report(ctx context.Context) error {
	result, err := ParseFile(r.spec.Report.Input)
	if err != nil {
		return errors.Wrapf(err, "error parse benchmark output")
	}
	req := &bhpb.GoBenchmarkReportResultRequest{
		JobId:     r.jobId,
		Package:   r.spec.Package,
		Results:   result,
		StartTime: r.startTime.UnixNano(),
		EndTime:   r.endTime.UnixNano(),
	}
	_, err = r.client.GoBenchmarkReportResult(ctx, req)
	return err
}
