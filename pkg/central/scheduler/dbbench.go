package scheduler

import (
	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

const (
	maxLoaderPack = 3 // at most 3 loader on a machine
)

var _ Scheduler = (*DbBench)(nil)

type DbBench struct {
	log *dlog.Logger
}

func NewDbBench() *DbBench {
	s := &DbBench{}
	dlog.NewStructLogger(log, s)
	return s
}

func (s *DbBench) AssignNode(nodes []pb.Node, specs []pb.NodeAssignmentSpec) ([]pb.AssignedNode, error) {
	if len(nodes) == 0 {
		return nil, errors.New("0 nodes available")
	}
	if len(nodes) < len(specs) {
		s.log.Warnf("only %d nodes available but requires %d nodes")
	}
	assignedNodes := make(map[string]*pb.AssignedNode, len(nodes))
	assignedSpecs := make(map[int]string, len(specs))

	// FIXME: only node.info.role is used
	// FIXME: node.role (current role) is ignored
	// FIXME: node state is ignored
	// FIXME: selector in spec is ignored
	for i, spec := range specs {
		for _, node := range nodes {
			// skipp assigned node
			if assignedNodes[node.Id] != nil {
				continue
			}
			// exact match of role or any
			if spec.Role == node.Info.Role || node.Info.Role == pb.Role_ANY {
				node.Role = spec.Role
				assignedNodes[node.Id] = &pb.AssignedNode{
					Node: node,
					Spec: spec,
				}
				assignedSpecs[i] = node.Id
			}
		}
	}

	// TODO: allow binpack loader into one node
	merr := errors.NewMultiErr()
	if len(assignedSpecs) != len(specs) {
		for i, spec := range specs {
			if assignedSpecs[i] == "" {
				merr.Append(errors.Errorf("spec name %s role %s not assigned", spec.Name, spec.Role))
			}
		}
	}

	res := make([]pb.AssignedNode, 0, len(assignedNodes))
	// the result is in the order of specification
	for i := 0; i < len(specs); i++ {
		res = append(res, *assignedNodes[assignedSpecs[i]])
	}
	return res, merr.ErrorOrNil()
}
