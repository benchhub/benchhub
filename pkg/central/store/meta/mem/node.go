package mem

import (
	"github.com/dyweb/gommon/errors"

	"github.com/benchhub/benchhub/pkg/central/store/meta"
	pbc "github.com/benchhub/benchhub/pkg/common/commonpb"
)

var _ meta.Provider = (*MetaStore)(nil)
var emptyNode = pbc.Node{}

// -- start of read --

func (s *MetaStore) NumNodes() (int, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.nodes), nil
}

func (s *MetaStore) FindNodeById(id string) (pbc.Node, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if n, ok := s.nodes[id]; ok {
		return n, nil
	} else {
		// TODO: might share a common error value, or use error code
		return emptyNode, errors.New("not found")
	}
}

func (s *MetaStore) ListNodes() ([]pbc.Node, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	nodes := make([]pbc.Node, 0, len(s.nodes))
	for id := range s.nodes {
		nodes = append(nodes, s.nodes[id])
	}
	return nodes, nil
}

func (s *MetaStore) ListNodesStatus() ([]pbc.NodeStatus, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	status := make([]pbc.NodeStatus, 0, len(s.status))
	for id := range s.status {
		status = append(status, s.status[id])
	}
	return status, nil
}

// -- end of read--

// -- start of write --

func (s *MetaStore) AddNode(id string, node pbc.Node) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.nodes[id]; ok {
		return errors.Errorf("node %s already exists", id)
	}
	s.nodes[id] = node
	return nil
}

func (s *MetaStore) UpdateNode(id string, node pbc.Node) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.nodes[id]; !ok {
		return errors.Errorf("node %s does not exists", id)
	}
	s.nodes[id] = node
	return nil
}

func (s *MetaStore) UpdateNodeStatus(id string, status pbc.NodeStatus) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.status[id] = status
	return nil
}

// -- end of write --

// -- start of delete --

func (s *MetaStore) RemoveNode(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.nodes[id]; !ok {
		return errors.Errorf("node %s does not exists", id)
	}
	delete(s.nodes, id)
	return nil
}

// -- end of delete --
