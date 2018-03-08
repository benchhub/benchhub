package job

import (
	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

const (
	maxLoaderPack = 3 // at most 3 loader on a machine
)

type Scheduler struct {
	log *dlog.Logger
}

// TODO: just 1:1 mapping for now, should have binpack of loader
type AssignedNode struct {
	Node pb.Node
	Spec pb.NodeAssignmentSpec
}

func NewScheduler() *Scheduler {
	s := &Scheduler{}
	dlog.NewStructLogger(log, s)
	return s
}

// AssignNode chose nodes based on spec
// TODO: only exact match is supported
func (s *Scheduler) AssignNode(nodes []pb.Node, specs []pb.NodeAssignmentSpec) ([]AssignedNode, error) {
	if len(nodes) == 0 {
		return nil, errors.New("0 nodes available")
	}
	if len(nodes) < len(specs) {
		s.log.Warnf("only %d nodes available but requires %d nodes")
	}
	assignedNodes := make(map[string]*AssignedNode, len(nodes))
	assignedSpecs := make(map[int]bool, len(specs))
	// FIXME: node state is ignored as well ....
	// FIXME: only role is used, other selector is ignored
	for i, spec := range specs {
		for _, node := range nodes {
			// skipp assigned node
			if assignedNodes[node.Id] != nil {
				continue
			}
			// exact match of role or any
			if spec.Role == node.Info.Role || node.Info.Role == pb.Role_ANY {
				if node.Info.Role == pb.Role_ANY {
					node.Role = spec.Role
					s.log.Debugf("update node %s from any to %s", node.Id, spec.Role)
				}
				assignedNodes[node.Id] = &AssignedNode{
					Node: node,
					Spec: spec,
				}
				assignedSpecs[i] = true
			}
		}
	}
	// TODO: allow pack loader into one node
	merr := errors.NewMultiErr()
	if len(assignedSpecs) != len(specs) {
		for i, spec := range specs {
			if !assignedSpecs[i] {
				merr.Append(errors.Errorf("spec name %s role %s not assigned", spec.Name, spec.Role))
			}
		}
	}
	// TODO: might use map[name]AssignedNode ...
	res := make([]AssignedNode, 0, len(assignedNodes))
	for id := range assignedNodes {
		//s.log.Info(assignedNodes[id].Spec.Name)
		res = append(res, *assignedNodes[id])
	}
	return res, merr.ErrorOrNil()
}
