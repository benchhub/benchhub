package generator

import (
	dlog "github.com/dyweb/gommon/log"
)

var (
	logReg = dlog.NewRegistry()
	log    = logReg.NewLogger()
)
