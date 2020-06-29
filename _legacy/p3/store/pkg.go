// Package store saves both meta data and raw time series
package store

import (
	dlog "github.com/dyweb/gommon/log"
)

var logReg = dlog.NewRegistry()
var log = logReg.Logger()
