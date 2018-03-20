package job

import (
	"sync"

	dlog "github.com/dyweb/gommon/log"
)

// Executor on local node execute plan received from central, one stage at a time
type Executor struct {
	mu sync.RWMutex

	log *dlog.Logger
}

func NewExecutor() *Executor {
	exc := &Executor{}
	dlog.NewStructLogger(log, exc)
	return exc
}
