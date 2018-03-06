package mem

import (
	"sync"

	dlog "github.com/dyweb/gommon/log"

	"github.com/benchhub/benchhub/pkg/central/store/meta"
	pbc "github.com/benchhub/benchhub/pkg/common/commonpb"
	"github.com/benchhub/benchhub/pkg/common/spec"
	"github.com/benchhub/benchhub/pkg/util/logutil"
)

var log = logutil.NewPackageLogger()

type MetaStore struct {
	mu sync.RWMutex

	nodes  map[string]pbc.Node
	status map[string]pbc.NodeStatus

	specs         map[string]spec.Job
	pendingSpecs  []string
	finishedSpecs []string

	log *dlog.Logger
}

func NewMetaStore() *MetaStore {
	s := &MetaStore{
		nodes:  make(map[string]pbc.Node, 10),
		status: make(map[string]pbc.NodeStatus, 10),
		specs:  make(map[string]spec.Job, 10),
	}
	dlog.NewStructLogger(log, s)
	return s
}

func init() {
	meta.RegisterProviderFactory("mem", func() meta.Provider {
		return NewMetaStore()
	})
}
