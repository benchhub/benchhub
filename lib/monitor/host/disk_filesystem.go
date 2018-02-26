package host

import (
	"syscall"
	"io"

	"github.com/dyweb/gommon/errors"
)

const (
	filesystemPath = "/"
)

var _ Stat = (*Filesystem)(nil)

type Filesystem struct {
	Type int64
	// Optimal transfer block size
	BlockSize uint64
	// Total data blocks in filesystem
	Blocks uint64
	// Free blocks in filesystem
	BlocksFree uint64
	// Free blocks available to unprivileged user
	BlocksAvail uint64

	// Total file nodes in filesystem
	Files uint64
	// Free file nodes in filesystem
	FilesFree uint64
	path      string
}

func NewFilesystem(path string) *Filesystem {
	if path == "" {
		path = filesystemPath
	}
	return &Filesystem{
		path: path,
	}
}

func (s *Filesystem) Path() string {
	if s.path == "" {
		return filesystemPath
	}
	return s.path
}

func (s *Filesystem) IsStatic() bool {
	return false
}

func (s *Filesystem) Update() error {
	return s.UpdateFrom(nil)
}

func (s *Filesystem) UpdateFrom(_ io.Reader) error {
	fs := syscall.Statfs_t{}
	// The  statfs() system call returns information about a mounted filesystem.
	// path is the pathname of any file within the mounted filesystem
	err := syscall.Statfs(s.Path(), &fs)
	if err != nil {
		return errors.Wrap(err, "error when syscall statfs")
	}
	// TODO: gopsutil has mapping for filesystem types
	s.Type = fs.Type
	s.BlockSize = uint64(fs.Bsize)
	s.Blocks = fs.Blocks
	s.BlocksFree = fs.Bfree
	s.BlocksAvail = fs.Bavail

	s.Files = fs.Files
	s.FilesFree = fs.Ffree
	// NOTE: Fsid is ignored, it is not major, minor number for filesystem in /proc/diskstats, from man statfs
	// The general idea is that f_fsid contains some random stuff such that the pair (f_fsid,ino) uniquely determines a file.
	return nil
}
