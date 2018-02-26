package host

import (
	"os"
	"runtime"
	"strings"

	"github.com/dyweb/gommon/errors"
)

// TODO: release information from lsb-release and os-release, saw it in gopsutil
// TODO:
// Machine contains information about physical node or vm, it won't change unless there are external forces
type Machine struct {
	NumCores       int
	MemTotal       uint64 // KB
	MemFree        uint64
	MemAvail       uint64 // NOTE: MemAvail = MemFree + Cache + Buffer
	DiskSpaceTotal uint64 // B
	DiskSpaceFree  uint64
	DiskInodeTotal uint64
	DiskInodeFree  uint64
	HostName       string
	BootTime       uint64

	mem Mem
	fs  Filesystem
}

func NewMachine() *Machine {
	return &Machine{}
}

func (s *Machine) IsStatic() bool {
	return true
}

func (s *Machine) Update() error {
	merr := errors.NewMultiErr()
	// TODO: NumCPU does not match /proc/stat on travis containerized build https://github.com/benchhub/benchhub/issues/9
	s.NumCores = runtime.NumCPU()
	hostname, err := os.Hostname()
	// TODO: merr might need AppendWrap
	merr.Append(errors.Wrap(err, "can't get host name"))
	s.HostName = hostname

	merr.Append(errors.Wrap(s.mem.Update(), "can't get host memory"))
	s.MemTotal = s.mem.MemTotal
	s.MemFree = s.mem.MemFree
	s.MemAvail = s.mem.MemAvailable

	merr.Append(errors.Wrap(s.fs.Update(), "can't get host disk space"))
	s.DiskSpaceTotal = s.fs.BlockSize * s.fs.Blocks
	// NOTE: we use BlocksAvail instead of BlocksFree because it's 'Free blocks available to unprivileged user'
	s.DiskSpaceFree = s.fs.BlockSize * s.fs.BlocksAvail
	s.DiskInodeTotal = s.fs.Files
	s.DiskInodeFree = s.fs.FilesFree

	// TODO: path should be configurable
	err = readFile("/proc/stat", func(line string) (stop bool) {
		parts := strings.Fields(line)
		if len(parts) == 0 {
			return false
		}
		if parts[0] == "btime" {
			s.BootTime = toUint64(parts[1])
			return true
		}
		return false
	})
	merr.Append(errors.Wrap(err, "can't get boot time"))

	return merr.ErrorOrNil()
}
