// Package server is BenhHub grpc server implementation.
package server

import (
	dlog "github.com/dyweb/gommon/log"
)

var logReg = dlog.NewRegistry()
var log = logReg.NewLogger()

// New returns a mega server that listens on specific address
func New(addr string) (*Mega, error) {
	return newMega(addr)
}
