package scheduler

import (
	pb "github.com/benchhub/benchhub/pkg/bhpb"
	"github.com/benchhub/benchhub/pkg/util/logutil"
)

var log = logutil.NewPackageLogger()

type Scheduler interface {
	AssignNode(nodes []pb.Node, specs []pb.NodeAssignmentSpec) ([]pb.AssignedNode, error)
}
