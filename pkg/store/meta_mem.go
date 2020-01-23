package store

import (
	"context"
	"sync"
	"time"

	"github.com/benchhub/benchhub/bhpb"
	"github.com/dyweb/gommon/errors"
	"github.com/gogo/protobuf/proto"
)

// meta_mem is in memory implementation for testing

type MetaMem struct {
	specMu     sync.RWMutex
	specNextId int
	specByHash map[string]int
	specs      []*bhpb.Spec

	jobMu     sync.RWMutex
	jobNextId int
	jobs      []*bhpb.Job
}

func NewMetaMem() *MetaMem {
	return &MetaMem{
		specNextId: 1,
		specByHash: make(map[string]int),
		jobNextId:  1,
	}
}

func (m *MetaMem) RegisterGoBenchmark(ctx context.Context, spec *bhpb.GoBenchmarkSpec) (*bhpb.JobRegisterResponse, error) {
	sid, err := m.registerSpec(spec)
	if err != nil {
		return nil, err
	}
	jid, err := m.registerJob(sid)
	if err != nil {
		return nil, err
	}
	return &bhpb.JobRegisterResponse{
		JobId:  int64(jid),
		SpecId: int64(sid),
	}, nil
}

func (m *MetaMem) registerSpec(spec *bhpb.GoBenchmarkSpec) (int, error) {
	h, err := HashGoBenchmarkSpec(spec)
	if err != nil {
		return -1, err
	}
	//log.Infof("hash is %s", h)

	m.specMu.Lock()
	defer m.specMu.Unlock()
	rowID, ok := m.specByHash[h]
	if ok {
		existing := m.specs[rowID-1]
		var decoded bhpb.GoBenchmarkSpec
		if err = proto.Unmarshal([]byte(existing.Payload), &decoded); err != nil {
			return -1, errors.Wrap(err, "error decode existing spec as go benchmark")
		}
		if proto.Equal(&decoded, spec) {
			return rowID, nil
		}
		// TODO: deal with hash collision
	}

	id := m.specNextId
	m.specNextId++
	b, err := proto.Marshal(spec)
	if err != nil {
		return -1, errors.Wrap(err, "error encode existing spec as go benchmark")
	}
	sspec := &bhpb.Spec{
		Id:            int64(id),
		BenchmarkType: bhpb.BenchmarkType_BENCHMARKTYPE_GO,
		Payload:       string(b),
		PayloadHash:   h,
		CreateTime:    time.Now().UnixNano(),
	}
	m.specs = append(m.specs, sspec)
	m.specByHash[h] = id
	return id, nil
}

func (m *MetaMem) registerJob(specId int) (int, error) {
	m.jobMu.Lock()
	defer m.jobMu.Unlock()

	id := m.jobNextId
	m.jobNextId++
	job := &bhpb.Job{
		Id:            int64(id),
		SpecId:        int64(specId),
		BenchmarkType: bhpb.BenchmarkType_BENCHMARKTYPE_GO,
		SubmitTime:    time.Now().UnixNano(),
	}
	m.jobs = append(m.jobs, job)
	return id, nil
}
