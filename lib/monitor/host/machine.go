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
	NumCores       int
	MemTotal       uint64
	MemFree        uint64
	MemAvail       uint64 // NOTE: MemAvail = MemFree + Cache + Buffer
	DiskSpaceTotal uint64
	DiskSpaceFree  uint64
	DiskInodeTotal uint64
	DiskInodeFree  uint64
	HostName       string

	mem Mem
	fs  Filesystem
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
		return errors.WithMessage(err, "can't get host memory")
	}
	s.MemTotal = s.mem.MemTotal
	s.MemFree = s.mem.MemFree
	s.MemAvail = s.mem.MemAvailable
	if err := s.fs.Update(); err != nil {
		return errors.WithMessage(err, "can't get host disk space")
	}
	s.DiskSpaceTotal = s.fs.BlockSize * s.fs.Blocks
	// NOTE: we use BlocksAvail instead of BlocksFree because it's 'Free blocks available to unprivileged user'
	s.DiskSpaceFree = s.fs.BlockSize * s.fs.BlocksAvail
	s.DiskInodeTotal = s.fs.Files
	s.DiskInodeFree = s.fs.FilesFree
	return nil
}
