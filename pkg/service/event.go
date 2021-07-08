package service

import (
	"context"
)

// Event ...
type Event interface {
	Start(ctx context.Context, r StartEvent) (err error)
	Finish(ctx context.Context, r FinishEvent) (err error)
}
