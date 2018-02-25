package host

import (
	"bufio"
	"bytes"
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
}

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

func toUint64(val string) uint64 {
	u, _ := strconv.ParseUint(val, 10, 64)
	return u
}
