package server

import (
	"sync"

	dlog "github.com/dyweb/gommon/log"

	pbc "github.com/benchhub/benchhub/pkg/bhpb"
)

// TODO: handle state of server

type StateMachine struct {
	mu      sync.RWMutex
	current pbc.NodeState
	log     *dlog.Logger
}

func NewStateMachine() (*StateMachine, error) {
	s := &StateMachine{
		current: pbc.NodeState_FINDING_CENTRAL,
	}
	dlog.NewStructLogger(log, s)
	return s, nil
}

func (s *StateMachine) Current() pbc.NodeState {
	s.mu.RLock()
	st := s.current
	s.mu.RUnlock()
	return st
}
func (s *StateMachine) RegisterSuccess() {
	s.mu.Lock()
	if s.current != pbc.NodeState_FINDING_CENTRAL {
		s.log.Warnf("previous state is not finding central but %s when register success", s.current)
	}
	s.updateState(pbc.NodeState_IDLE)
	s.mu.Unlock()
}

func (s *StateMachine) updateState(ns pbc.NodeState) {
	s.log.Infof("update state from %s to %s", s.current, ns)
	s.current = ns
}
