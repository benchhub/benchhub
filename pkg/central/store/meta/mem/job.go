package mem

import (
	pb "github.com/benchhub/benchhub/pkg/bhpb"

	"github.com/dyweb/gommon/errors"
	"strconv"
)

var emptySpec = pb.JobSpec{}

func (s *MetaStore) GetJobSpec(id string) (pb.JobSpec, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if job, ok := s.specs[id]; ok {
		return job, nil
	} else {
		return emptySpec, errors.New("not found")
	}
}

func (s *MetaStore) AddJobSpec(job pb.JobSpec) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.jobCounter++
	id := strconv.Itoa(s.jobCounter)
	s.specs[id] = job
	s.pendingSpecs = append(s.pendingSpecs, id)
	return id, nil
}

func (s *MetaStore) GetPendingJob() (job pb.JobSpec, empty bool, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	// FIFO
	if len(s.pendingSpecs) == 0 {
		return emptySpec, true, nil
	}
	// TODO: remove job?
	job = s.specs[s.pendingSpecs[0]]
	s.pendingSpecs = s.pendingSpecs[1:]
	return job, false, nil
}

func (s *MetaStore) PushbackJob(id string, job pb.JobSpec) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.specs[id] = job
	s.pendingSpecs = append(s.pendingSpecs, id)
	return nil
}
