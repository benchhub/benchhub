package mem

import (
	"sync"

	dlog "github.com/dyweb/gommon/log"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
	"github.com/benchhub/benchhub/pkg/central/store/meta"
	"github.com/benchhub/benchhub/pkg/util/logutil"
)

var log = logutil.NewPackageLogger()

type MetaStore struct {
	mu sync.RWMutex

	nodes  map[string]pb.Node
	status map[string]pb.NodeStatus

	specs         map[string]pb.JobSpec
	pendingSpecs  []string
	finishedSpecs []string

	log *dlog.Logger
}

func NewMetaStore() *MetaStore {
	s := &MetaStore{
		nodes:  make(map[string]pb.Node, 10),
		status: make(map[string]pb.NodeStatus, 10),
		specs:  make(map[string]pb.JobSpec, 10),
	}
	dlog.NewStructLogger(log, s)
	return s
}

func init() {
	meta.RegisterProviderFactory("mem", func() meta.Provider {
		return NewMetaStore()
	})
}
