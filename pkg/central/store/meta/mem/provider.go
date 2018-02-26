package mem

import (
	"sync"

	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"

	"github.com/benchhub/benchhub/pkg/central/store/meta"
	pbc "github.com/benchhub/benchhub/pkg/common/commonpb"
)

var _ meta.Provider = (*MetaStore)(nil)
var emptyNode = pbc.Node{}

type MetaStore struct {
	mu sync.RWMutex

	nodes map[string]pbc.Node
	log   *dlog.Logger
}

func NewMetaStore() *MetaStore {
	s := &MetaStore{
		nodes: make(map[string]pbc.Node, 10),
	}
	dlog.NewStructLogger(log, s)
	return s
}

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
