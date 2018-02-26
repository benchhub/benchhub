package host

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/dyweb/gommon/errors"

	"github.com/benchhub/benchhub/lib/monitor/util/logutil"
)

var log = logutil.NewPackageLogger()

type ProcFile interface {
	Path() string
}

type Stat interface {
	IsStatic() bool
	Update() error
	UpdateFrom(r io.Reader) error
}

// TODO: Update that accepts a reader (also makes test easier ...), and we can re use opened file like /proc/stat

type lineHandler func(line string) (stop bool)

// based on gosigar sigar_linux(_common).go
func readFile(path string, cb lineHandler) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrapf(err, "can't read file %s", path)
	}
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		if cb(scanner.Text()) {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return errors.Wrapf(err, "error when scan file %s", path)
	}
	return nil
}

func readFrom(r io.Reader, cb lineHandler) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if cb(scanner.Text()) {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return errors.Wrap(err, "error when scan io.Reader")
	}
	return nil
}

func toUint64(val string) uint64 {
	u, _ := strconv.ParseUint(val, 10, 64)
	return u
}
