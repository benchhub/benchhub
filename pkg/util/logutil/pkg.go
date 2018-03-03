package logutil

import (
	goicelog "github.com/at15/go.ice/ice/util/logutil"
	"github.com/dyweb/gommon/log"

	libmon "github.com/benchhub/benchhub/lib/monitor/util/logutil"
	librunner "github.com/benchhub/benchhub/lib/runner/util/logutil"
)

var Registry = log.NewApplicationLogger()

func NewPackageLogger() *log.Logger {
	l := log.NewPackageLoggerWithSkip(1)
	Registry.AddChild(l)
	return l
}

func init() {
	Registry.AddChild(goicelog.Registry)
	Registry.AddChild(libmon.Registry)
	Registry.AddChild(librunner.Registry)
}
