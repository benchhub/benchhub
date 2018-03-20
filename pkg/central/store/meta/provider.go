package meta

import (
	"sort"
	"sync"

	"github.com/dyweb/gommon/errors"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

var (
	providerMu        sync.Mutex
	providerFactories = make(map[string]ProviderFactory)
)

type NodeProvider interface {
	// read
	NumNodes() (int, error)
	// TODO: special error for not found?
	// NOTE: we always return by value to avoid (my) common mistake of pointer pointing to last element in for .. range
	FindNodeById(id string) (pb.Node, error)
	ListNodes() ([]pb.Node, error)
	ListNodesStatus() ([]pb.NodeStatus, error)

	// write
	AddNode(id string, node pb.Node) error
	UpdateNode(id string, node pb.Node) error
	UpdateNodeStatus(id string, status pb.NodeStatus) error

	// delete
	RemoveNode(id string) error
}

type JobProvider interface {
	// TODO: how to avoid a user submit same job twice?
	AddJobSpec(job pb.JobSpec) (id string, err error)
	GetJobSpec(id string) (pb.JobSpec, error)
	// GetPending job get a spec from start of a queue
	GetPendingJob() (job pb.JobSpec, empty bool, err error)
	// PushbackJob put a spec to the end of queue after it is retrieved using GetPendingJob TODO: put to head of queue instead?
	PushbackJob(id string, job pb.JobSpec) error
	// TODO: API for both spec and job status ...
}

type Provider interface {
	NodeProvider
	JobProvider
}

// TODO: factory should accept config, this is needed for rdbms
type ProviderFactory func() Provider

func GetProvider(name string) (Provider, error) {
	providerMu.Lock()
	defer providerMu.Unlock()
	if f, ok := providerFactories[name]; !ok {
		return nil, errors.Errorf("provider %s is not registered", name)
	} else {
		// TODO: factory may also return error once we have config
		return f(), nil
	}
}

func RegisterProviderFactory(name string, factory ProviderFactory) {
	providerMu.Lock()
	defer providerMu.Unlock()
	if _, dup := providerFactories[name]; dup {
		log.Panicf("RegisterProviderFactory is called twice for %s", name)
	}
	providerFactories[name] = factory
	// FIXED: this logger never showed up ... because this function is called before the cli application set the
	log.Debugf("register provider factory %s", name)
}

func Providers() []string {
	providerMu.Lock()
	defer providerMu.Unlock()
	list := make([]string, 0, len(providerFactories))
	for name := range providerFactories {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}
