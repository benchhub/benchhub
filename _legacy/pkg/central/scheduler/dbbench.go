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
	// we can put multiple loader in one node, so it may not be an error
	if len(nodes) < len(specs) {
		s.log.Warnf("only %d nodes available but requires %d nodes", len(nodes), len(specs))
	}

	merr := errors.NewMultiErr()
	assignedNodes := make(map[string]*pb.AssignedNode, len(nodes))
	assignedSpecs := make(map[int]string, len(specs))

	// FIXME: only node.info.role is used
	// FIXME: node.role (current role) is ignored
	// FIXME: node state is ignored
	// FIXME: selector in spec is ignored
NextSpec:
	for i, spec := range specs {
		s.log.Infof("spec %s role %s", spec.Properties.Name, spec.Properties.Role)
	NextNode:
		for _, node := range nodes {
			// skipp assigned node
			if assignedNodes[node.Info.Id] != nil {
				continue NextNode
			}
			var (
				success pb.ScheduleSuccess
				err     error
			)
			switch spec.Properties.Role {
			case pb.Role_DATABASE:
				success, err = s.scheduleDb(node, spec)
			case pb.Role_LOADER:
				success, err = s.scheduleLoader(node, spec)
			default:
				merr.Append(errors.Errorf("invalid spec unknown role %s for %s", spec.Properties.Role, spec.Properties.Name))
			}
			// exact match of role or any
			if spec.Properties.Role == node.Info.Config.Role ||
				node.Info.Config.Role == pb.Role_ANY {
				s.log.Infof("chose node %s for spec %s", node.Info.Id, spec.Properties.Name)
				assignedNodes[node.Info.Id] = &pb.AssignedNode{
					Node: node.Info,
					Spec: spec,
				}
				assignedSpecs[i] = node.Info.Id
				continue NextSpec
			}
		}
	}

	//s.log.Info("assign finished")

	// TODO: allow binpack loader into one node
	if len(assignedSpecs) != len(specs) {
		for i, spec := range specs {
			if assignedSpecs[i] == "" {
				merr.Append(errors.Errorf("spec name %s role %s not assigned", spec.Properties.Name, spec.Properties.Role))
			}
		}
		return nil, merr.ErrorOrNil()
	}

	//s.log.Info("combine assign result")

	res := make([]pb.AssignedNode, 0, len(assignedNodes))
	// the result is in the order of specification
	// FIXME: this would cause panic: runtime error: invalid memory address or nil pointer dereference
	// because not all specs are assigned
	for i := 0; i < len(specs); i++ {
		res = append(res, *assignedNodes[assignedSpecs[i]])
	}
	return res, merr.ErrorOrNil()
}

// TODO: reason for why nodes are filtered out
func (s *DbBench) filterDb(nodes []pb.Node) ([]pb.Node, error) {

}

func (s *DbBench) scheduleDb(node pb.Node, spec pb.NodeAssignmentSpec) (pb.ScheduleSuccess, error) {
	if spec.Properties.Role != pb.Role_DATABASE {
		return pb.ScheduleSuccess_UNKNOWN_SCHEDULE_SUCCESS,
			errors.Errorf("incompatible role %s passed for database", spec.Properties.Role)
	}
	if node.Status.State == pb.NodeState_NODE_IDLE {
		if node.Info.Config.Role == pb.Role_DATABASE {
			return pb.ScheduleSuccess_IDLE_EXACT_ROLE_MATCH, nil
		}
		if node.Info.Config.Role == pb.Role_ANY {
			return pb.ScheduleSuccess_IDLE_ANY_ROLE_MATCH, nil
		}
		return pb.ScheduleSuccess_UNKNOWN_SCHEDULE_SUCCESS,
			errors.Wrap(&Error{Code: pb.ScheduleFail_IDLE_ROLE_NO_MATCH}, "failed to schedule db")
	}
	// TODO: check if this node is already running database
	return pb.ScheduleSuccess_UNKNOWN_SCHEDULE_SUCCESS, errors.New("not implemented, schedule on non idle node")
}

func (s *DbBench) scheduleLoader(node pb.Node, spec pb.NodeAssignmentSpec) (pb.ScheduleSuccess, error) {
	if spec.Properties.Role != pb.Role_LOADER {
		return errors.Errorf("in compatible role %s passed to dbBest", spec.Properties.Role)
	}
}
