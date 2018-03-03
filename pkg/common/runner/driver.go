package runner

import "context"

type Driver interface {
	Run(ctx context.Context) error
}
