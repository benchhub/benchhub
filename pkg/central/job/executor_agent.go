package job

import (
	"sync"

	dlog "github.com/dyweb/gommon/log"
)

var _ Executor = (*AgentExecutor)(nil)

// AgentExecutor dispatch plan to node agent on worker nodes
type AgentExecutor struct {
	*sync.RWMutex

	log *dlog.Logger
}

func NewAgentExecutor() *AgentExecutor {
	exc := &AgentExecutor{}
	dlog.NewStructLogger(log, exc)
	return exc
}

//func (exc *AgentExecutor) RLock() {
//	exc.mu.RLock()
//}
//
//func (exc *AgentExecutor) RUnlock() {
//	exc.mu.RUnlock()
//}
//
//func (exc *AgentExecutor) Lock() {
//	exc.mu.Lock()
//}
//
//func (exc *AgentExecutor) Unlock() {
//	exc.mu.Unlock()
//}
