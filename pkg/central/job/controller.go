package job

import (
	dlog "github.com/dyweb/gommon/log"
)

type Controller struct {
	scheduler *Scheduler
	log       *dlog.Logger
}

func NewController() *Controller {
	c := &Controller{
		scheduler: NewScheduler(),
	}
	return c
}
