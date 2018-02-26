package host

import (
	"io"
	"runtime"
	"strings"
)

// TODO: /proc/stat has more than just cpu
// TODO: cadvisor change log says uint64 could overflow in version 0.23.4
const (
	// TODO: it seems it is always the case for x86, sysconf(_SC_CLK_TCK) would have the right value
	// TODO: should we / CPUUserHz like we did in xephon-k collector
	CPUUserHz   = 100
	cpuStatPath = "/proc/stat"
)

var _ Stat = (*Cpus)(nil)

type Cpus struct {
	Total Cpu
	Cores []Cpu
	path  string
	num   int
}

type Cpu struct {
	User      uint64
	Nice      uint64
	System    uint64
	Idle      uint64
	IOWait    uint64
	Irq       uint64
	SoftIrq   uint64
	Steal     uint64
	Guest     uint64
	GuestNice uint64
}

func NewCpus(path string) *Cpus {
	if path == "" {
		path = cpuStatPath
	}
	return &Cpus{
		path: path,
		num:  runtime.NumCPU(),
	}
}

func (s *Cpus) Path() string {
	return s.path
}

func (s *Cpus) IsStatic() bool {
	return false
}

func (s *Cpus) UpdateFrom(r io.Reader) error {
	cores := make([]Cpu, 0, s.num)
	err := readFrom(r, func(line string) (stop bool) {
		parts := strings.Fields(line)
		if len(parts) == 0 {
			return false
		}
		head := parts[0]
		if head[:3] != "cpu" {
			return true
		}
		core := Cpu{}
		parseCpuStatLine(parts[1:], &core)
		if head == "cpu" {
			s.Total = core
		} else {
			cores = append(cores, core)
		}
		return false
	})
	s.Cores = cores
	return err
}

// The amount of time, measured in units of USER_HZ (1/100ths of a second on most architectures,
// use sysconf(_SC_CLK_TCK) to obtain the right value)
func parseCpuStatLine(fields []string, cpu *Cpu) {
	cpu.User = toUint64(fields[0])
	cpu.Nice = toUint64(fields[1])
	cpu.System = toUint64(fields[2])
	cpu.Idle = toUint64(fields[3])
	cpu.IOWait = toUint64(fields[4])
	cpu.Irq = toUint64(fields[5])
	cpu.SoftIrq = toUint64(fields[6])
	cpu.Steal = toUint64(fields[7])
	// TODO: does vm have this?
	cpu.Guest = toUint64(fields[8])
	cpu.GuestNice = toUint64(fields[9])
}
