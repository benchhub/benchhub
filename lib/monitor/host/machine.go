package host

import (
	"os"
	"runtime"

	"github.com/pkg/errors"
)

// TODO: release information from lsb-release and os-release, saw it in gopsutil
// TODO:
// Machine contains information about physical node or vm, it won't change unless there are external forces
type Machine struct {
	NumCores int
	Mem      uint64
	HostName string

	mem Mem
}

func (s *Machine) IsStatic() bool {
	return true
}

func (s *Machine) Update() error {
	// TODO: NumCPU does not match /proc/stat on travis containerized build https://github.com/benchhub/benchhub/issues/9
	s.NumCores = runtime.NumCPU()
	hostname, err := os.Hostname()
	if err != nil {
		return errors.Wrap(err, "can't get host name")
	}
	s.HostName = hostname
	if err := s.mem.Update(); err != nil {
		return err
	}
	s.Mem = s.mem.MemTotal
	return nil
}
