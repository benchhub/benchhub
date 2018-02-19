package host

import (
	"strings"
)

const (
	diskStatPath      = "/proc/diskstats"
	diskStatNumFileds = 14
	// TODO: /proc/partitions
	// TODO: /proc
)

// https://www.kernel.org/doc/Documentation/ABI/testing/procfs-diskstats
//The /proc/diskstats file displays the I/O statistics
//of block devices. Each line contains the following 14
//fields:
//1 - major number
//2 - minor mumber
//3 - device name
//4 - reads completed successfully
//5 - reads merged
//6 - sectors read
//7 - time spent reading (ms)
//8 - writes completed
//9 - writes merged
//10 - sectors written
//11 - time spent writing (ms)
//12 - I/Os currently in progress
//13 - time spent doing I/Os (ms)
//14 - weighted time spent doing I/Os (ms)
//For more details refer to Documentation/iostats.txt
// https://www.kernel.org/doc/Documentation/iostats.txt

// https://github.com/salesforce/LinuxTelemetry/blob/master/plugins/diskstats.py
// TODO: what are loop device? https://en.wikipedia.org/wiki/Loop_device

type BlockDevices struct {
	Devices map[string]*BlockDevice
	path    string
}

type BlockDevice struct {
	// /dev/block/major:minor
	Major uint64
	Minor uint64
	Name  string

	// Field 1 - 11
	ReadsCompleted uint64
	// Reads and writes which are adjacent to each other may be merged for
	// efficiency.  Thus two 4K reads may become one 8K read before it is
	// ultimately handed to the disk, and so it will be counted (and queued)
	// as only one I/O.  This field lets you know how often this was done
	ReadsMerged        uint64
	SectorsRead        uint64 // TODO: sector size?
	TimeSpentReadingMs uint64
	WritesCompleted    uint64
	WritesMerged       uint64
	SectorsWritten     uint64 // TODO: same as read
	TimeSpentWritingMs uint64
	// Field 9: The only field that should go to zero. Incremented as requests are given to appropriate struct request_queue and decremented as they finish
	IoInProgress uint64
	// Field 10: This field increases so long as field 9 is nonzero
	TimeSpentIoMs         uint64
	WeightedTimeSpentIoMs uint64
}

func NewBlockDevices(path string) *BlockDevices {
	if path == "" {
		path = diskStatPath
	}
	return &BlockDevices{
		Devices: make(map[string]*BlockDevice, 5),
		path:    path,
	}
}

func (s *BlockDevices) Path() string {
	return s.path
}

func (s *BlockDevices) Update() error {
	devices := make(map[string]*BlockDevice, len(s.Devices))
	err := readFile(s.Path(), func(line string) (stop bool) {
		parts := strings.Fields(line)
		// invalid, but don't stop
		if len(parts) < 14 {
			return false
		}
		device := &BlockDevice{}
		parseDiskStatLine(parts, device)
		devices[device.Name] = device
		return false
	})
	s.Devices = devices
	return err
}

func parseDiskStatLine(fields []string, device *BlockDevice) {
	device.Major = toUint64(fields[0])
	device.Minor = toUint64(fields[1])
	device.Name = fields[2]

	device.ReadsCompleted = toUint64(fields[3])
	device.ReadsMerged = toUint64(fields[4])
	device.SectorsRead = toUint64(fields[5])
	device.TimeSpentReadingMs = toUint64(fields[6])

	device.WritesCompleted = toUint64(fields[7])
	device.WritesMerged = toUint64(fields[8])
	device.SectorsWritten = toUint64(fields[9])
	device.TimeSpentWritingMs = toUint64(fields[10])

	device.IoInProgress = toUint64(fields[11])
	device.TimeSpentIoMs = toUint64(fields[12])
	device.WeightedTimeSpentIoMs = toUint64(fields[13])
}
