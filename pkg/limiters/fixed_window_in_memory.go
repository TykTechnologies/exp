package limiters

import (
	"context"
	"sync"
	"time"
)

// FixedWindowInMemory is an in-memory implementation of FixedWindowIncrementer.
type FixedWindowInMemory struct {
	mu     sync.Mutex
	c      int64
	window time.Time
}

// NewFixedWindowInMemory creates a new instance of FixedWindowInMemory.
func NewFixedWindowInMemory() *FixedWindowInMemory {
	return &FixedWindowInMemory{}
}

// Increment increments the window's counter.
func (f *FixedWindowInMemory) Increment(ctx context.Context, window time.Time, _ time.Duration) (int64, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if window != f.window {
		f.c = 0
		f.window = window
	}
	f.c++
	return f.c, ctx.Err()
}
