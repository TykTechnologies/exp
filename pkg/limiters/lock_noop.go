package limiters

import (
	"context"
)

// LockNoop is a no-op implementation of the DistLocker interface.
// It should only be used with the in-memory backends as they are already thread-safe and don't need distributed locks.
type LockNoop struct {
}

// NewLockNoop creates a new LockNoop.
func NewLockNoop() *LockNoop {
	return &LockNoop{}
}

// Lock imitates locking.
func (n LockNoop) Lock(ctx context.Context) error {
	return ctx.Err()
}

// Unlock does nothing.
func (n LockNoop) Unlock(_ context.Context) error {
	return nil
}
