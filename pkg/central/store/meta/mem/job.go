package mem

import (
	"github.com/benchhub/benchhub/pkg/common/spec"
	"github.com/dyweb/gommon/errors"
)

var emptySpec = spec.Job{}

func (s *MetaStore) GetJobSpec(id string) (spec.Job, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if job, ok := s.specs[id]; ok {
		return job, nil
	} else {
		return emptySpec, errors.New("not found")
	}
}

func (s *MetaStore) AddJobSpec(id string, job spec.Job) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.specs[id]; ok {
		return errors.Errorf("job %s already exists", id)
	}
	s.specs[id] = job
	s.pendingSpecs = append(s.pendingSpecs, id)
	return nil
}

func (s *MetaStore) GetPendingJob() (job spec.Job, empty bool, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	// FIFO
	if len(s.pendingSpecs) == 0 {
		return emptySpec, true, nil
	}
	job = s.specs[s.pendingSpecs[0]]
	s.pendingSpecs = s.pendingSpecs[1:]
	return job, false, nil
}
