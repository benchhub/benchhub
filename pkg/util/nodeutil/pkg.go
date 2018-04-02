// Package nodeutil provides util for machine (vm) info, i.e. uid, net addr
package nodeutil

import (
	"time"

	"github.com/benchhub/benchhub/pkg/util/logutil"
)

// TODO
//
// get node ip address https://github.com/benchhub/benchhub/issues/18

var log = logutil.NewPackageLogger()

var (
	startTime time.Time
)

// StartTime is the process start time
func StartTime() time.Time {
	return startTime
}

func init() {
	startTime = time.Now()
}
