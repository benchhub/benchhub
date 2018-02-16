package host

import (
	"os"
	"runtime"

	"github.com/pkg/errors"
)

// Machine contains information about physical node or vm, it won't change unless there are external forces
type Machine struct {
	NumCores int
	HostName string
}

func (s *Machine) IsStatic() bool {
	return true
}

func (s *Machine) Update() error {
	s.NumCores = runtime.NumCPU()
	hostname, err := os.Hostname()
	if err != nil {
		return errors.Wrap(err, "can't get host name")
	}
	s.HostName = hostname
	return nil
}
