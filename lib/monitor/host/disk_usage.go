package host

// disk usage is basically counting '/' or all the mounted filesystems ...
// man statfs
// https://gist.github.com/sparrc/b1b4068141d192f95e54
// https://gist.github.com/lunny/9828326
//const (
//	diskUsagePath = "/"
//)
//
//type DiskUsage struct {
//	path string
//}
//
//func NewDiskUsage(path string) *DiskUsage {
//	if path == "" {
//		path = diskUsagePath
//	}
//	return &DiskUsage{path: path}
//}
//
//func (s *DiskUsage) Update() error {
//	// The  statfs() system call returns information about a mounted filesystem.
//	// path is the pathname of any file within the mounted filesystem
//	fs := syscall.Statfs_t{}
//	err := syscall.Statfs(s.path, &fs)
//	if err != nil {
//		return errors.Wrap(err, "error when syscall statfs")
//	}
//
//}
