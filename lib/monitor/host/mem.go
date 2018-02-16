package host

import "strings"

const (
	memStatPath = "/proc/meminfo"
)

// TODO: the code for reading mem should be generated, also we can add descriptions
type Mem struct {
	MemTotal     uint64
	MemFree      uint64
	MemAvailable uint64
	Buffers      uint64
	Cached       uint64
	Active       uint64
	Inactive     uint64
	Dirty        uint64
	Writeback    uint64 // Memory which is actively being written back to the disk
	Mapped       uint64 // Files which have been mapped into memory (with mmap(2)), such as librarie
	Shmem        uint64 // Amount of memory consumed in tmpfs(5) filesystems
	Slab         uint64 // In-kernel data structures cache
	SReclaimable uint64
	SUnreclaim   uint64
	KernelStack  uint64 // Amount of memory allocated to kernel stacks
	PageTables   uint64
	WritebackTmp uint64 // Memory used by FUSE for temporary writeback buffers
	HugePagesize uint64
	DirectMap4k  uint64
	DirectMap2M  uint64
	DirectMap1G  uint64

	SwapCached uint64
	SwapTotal  uint64
	SwapFree   uint64

	path string
}

func NewMem(path string) *Mem {
	if path == "" {
		path = memStatPath
	}
	return &Mem{
		path: path,
	}
}

func (s *Mem) Path() string {
	return s.path
}

func (s *Mem) Update() error {
	err := readFile(s.Path(), func(line string) (stop bool) {
		parts := strings.Fields(line)
		if len(parts) < 2 {
			return false
		}
		// omit the `:`, i.e. `MemTotal:` -> `MemTotal`
		head := parts[0][0 : len(parts[0])-1]
		value := toUint64(parts[1])
		switch head {
		case "MemTotal":
			s.MemTotal = value
		case "MemFree":
			s.MemFree = value
		case "MemAvailable":
			s.MemAvailable = value
		case "Buffers":
			s.Buffers = value
		case "Cached":
			s.Cached = value
		case "Active":
			s.Active = value
		case "Inactive":
			s.Inactive = value
		case "Dirty":
			s.Dirty = value
		case "Writeback":
			s.Writeback = value
		case "Mapped":
			s.Mapped = value
		case "Shmem":
			s.Shmem = value
		case "Slab":
			s.Slab = value
		case "SReclaimable":
			s.SReclaimable = value
		case "SUnreclaim":
			s.SUnreclaim = value
		case "KernelStack":
			s.KernelStack = value
		case "PageTables":
			s.PageTables = value
		case "WritebackTmp":
			s.WritebackTmp = value
		case "HugePagesize":
			s.HugePagesize = value
		case "DirectMap4k":
			s.DirectMap4k = value
		case "DirectMap2M":
			s.DirectMap2M = value
		case "DirectMap1G":
			s.DirectMap1G = value
			// Swap
		case "SwapCached":
			s.SwapCached = value
		case "SwapTotal":
			s.SwapTotal = value
		case "SwapFree":
			s.SwapFree = value
		default:
			// do nothing
			//fmt.Println(head)
		}
		return false
	})
	return err
}
