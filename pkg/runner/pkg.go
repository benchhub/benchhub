// Package runner run shell command or docker defined in Task
// But driver (executor) does not know if this a long running command, it's defined in the task
package runner

import (
	"context"
	"github.com/benchhub/benchhub/pkg/util/logutil"
)

var log = logutil.NewPackageLogger()

//type RunWithContext interface {
//	RunWithContext(ctx context.Context) error
//}
