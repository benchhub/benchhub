package scheduler

import (
	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

// Error is reason for failed to schedule a spec to one node
type Error struct {
	Code pb.ScheduleFail
}

func (e *Error) Error() string {
	return e.Code.String()
}
