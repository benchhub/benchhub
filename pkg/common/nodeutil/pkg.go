// Package nodeutil provides util for machine (vm) info, i.e. uid, net addr
package nodeutil

import (
	"github.com/benchhub/benchhub/pkg/util/logutil"
	"time"
)

// TODO
//
// get node capacity, use lib/monitor
// get node ip address https://github.com/benchhub/benchhub/issues/18

var log = logutil.NewPackageLogger()
var startTime = time.Now()

// StartTime is the process start time
func StartTime() time.Time {
	return startTime
}
