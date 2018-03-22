package job

import (
	dlog "github.com/dyweb/gommon/log"
)

type Manager struct {
	executor *Executor
	log      *dlog.Logger
}

func NewManager() *Manager {
	c := &Manager{
		executor: NewExecutor(),
	}
	dlog.NewStructLogger(log, c)
	return c
}
