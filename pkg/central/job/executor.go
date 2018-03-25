package job

import (
	"sync"

	dlog "github.com/dyweb/gommon/log"
)

// Executor runs plan on a single node, coordination between nodes is done by manager
// - simply print command out (for dry run)
// - run everything locally, assume the job spec can be run locally, i.e. everything is localhost
// - dispatch job to remote agent
type Executor interface {
	rwMutex
}

type rwMutex interface {
	RLock()
	RUnlock()
	Lock()
	Unlock()
}

var _ Executor = (*TestExecutor)(nil)
var _ Executor = (*EchoExecutor)(nil)
var _ Executor = (*LocalExecutor)(nil)

type TestExecutor struct {
	*sync.RWMutex

	log *dlog.Logger
}

type EchoExecutor struct {
	*sync.RWMutex

	log *dlog.Logger
}

type LocalExecutor struct {
	*sync.RWMutex

	log *dlog.Logger
}
