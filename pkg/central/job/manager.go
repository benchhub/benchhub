package job

import (
	dlog "github.com/dyweb/gommon/log"
)

type Manager struct {
	planner  *Planner
	executor *Executor
	log      *dlog.Logger
}

func NewManager() *Manager {
	c := &Manager{
		planner:  NewPlanner(),
		executor: NewExecutor(),
	}
	dlog.NewStructLogger(log, c)
	return c
}
