package runner

import (
	"context"
	"io"
	"os"
	"os/exec"
	"syscall"

	"github.com/dyweb/gommon/errors"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

type Exec struct {
	spec   pb.ExecSpec
	stdout io.Writer
	stderr io.Writer
}

// NewExec use os.Stdout and os.Stderr as default redirect
// TODO: should return error and creates the cmd here, for shell, split shell quote is also needed
func NewExec(s pb.ExecSpec) *Exec {
	return &Exec{spec: s, stdout: os.Stdout, stderr: os.Stderr}
}

// Run wait for the command to finish or kill the process group when context is canceled
func (e *Exec) Run(ctx context.Context) error {
	// NOTE: we are not using exec.CommandExec because it only kill the process, we may need to kill process group
	cmd := exec.Command(e.spec.Command, e.spec.Args...)
	cmd.Stdout = e.stdout
	cmd.Stderr = e.stderr
	// TODO: we need this use KillProcessGroup, but what does setpgid really do?
	cmd.SysProcAttr = SetProcessGroup(cmd.SysProcAttr)
	if err := cmd.Start(); err != nil {
		return errors.Wrap(err, "failed to start")
	}
	waitCh := make(chan error)
	go func() {
		err := cmd.Wait()
		if err != nil {
			err = errors.Wrap(err, "error when wait process to finish")
		}
		waitCh <- err
	}()
	select {
	case err := <-waitCh:
		return err
	case <-ctx.Done():
		// TODO: ctx.Err()?
		return KillProcessGroup(cmd.Process)
	}
}

// SetProcessGroup is copied from gitlab-runner helpers process_group_unix
// without calling this KillProcessGroup won't work
// TODO: I guess it allows you to kill all the child processes of a processes?
func SetProcessGroup(attr *syscall.SysProcAttr) *syscall.SysProcAttr {
	if attr == nil {
		return &syscall.SysProcAttr{Setpgid: true}
	}
	attr.Setpgid = true
	return attr
}

// KillProcessGroup is copied gitlab-runner helpers process_group_unix
// man 2 kill
func KillProcessGroup(process *os.Process) error {
	if process == nil {
		return errors.New("nil *os.Process passed")
	}
	if process.Pid > 0 {
		if err := syscall.Kill(-process.Pid, syscall.SIGKILL); err != nil {
			return errors.Wrapf(err, "failed to kill process group %d", process.Pid)
		}
	} else {
		if err := process.Kill(); err != nil {
			return errors.Wrap(err, "failed to kill process alone")
		}
	}
	return nil
}
