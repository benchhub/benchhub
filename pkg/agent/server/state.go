package server

import (
	"sync"

	dlog "github.com/dyweb/gommon/log"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

// TODO: handle state of server

type StateMachine struct {
	mu      sync.RWMutex
	current pb.NodeState
	log     *dlog.Logger
}

func NewStateMachine() (*StateMachine, error) {
	s := &StateMachine{
		current: pb.NodeState_NODE_FINDING_CENTRAL,
	}
	dlog.NewStructLogger(log, s)
	return s, nil
}

func (s *StateMachine) Current() pb.NodeState {
	s.mu.RLock()
	st := s.current
	s.mu.RUnlock()
	return st
}

func (s *StateMachine) RegisterSuccess() {
	s.mu.Lock()
	if s.current != pb.NodeState_NODE_FINDING_CENTRAL {
		s.log.Warnf("previous state is not finding central but %s when register success", s.current)
	}
	s.updateState(pb.NodeState_NODE_IDLE)
	s.mu.Unlock()
}

func (s *StateMachine) HeartbeatFailed() {
	s.mu.Lock()
	if s.current == pb.NodeState_NODE_FINDING_CENTRAL {
		s.log.Warnf("previous state is finding central but heart beat failed is called", s.current)
	}
	s.updateState(pb.NodeState_NODE_FINDING_CENTRAL)
	s.mu.Unlock()
}

func (s *StateMachine) updateState(ns pb.NodeState) {
	s.log.Infof("update state from %s to %s", s.current, ns)
	s.current = ns
}
