package lock

import (
	"context"
)

// Lock define common func
type Lock interface {
	Lock(ctx context.Context) (bool, error)
	UnLock(ctx context.Context) (bool, error)
}
