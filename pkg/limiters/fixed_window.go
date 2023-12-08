package limiters

import (
	"context"
	"sync"
	"time"
)

// FixedWindow implements a Fixed Window rate limiting algorithm.
//
// Simple and memory efficient algorithm that does not need a distributed lock.
// However it may be lenient when there are many requests around the boundary between 2 adjacent windows.
type FixedWindow struct {
	backend  FixedWindowIncrementer
	clock    Clock
	rate     time.Duration
	capacity int64
	mu       sync.Mutex
	window   time.Time
	overflow bool
}

// NewFixedWindow creates a new instance of FixedWindow.
// Capacity is the maximum amount of requests allowed per window.
// Rate is the window size.
func NewFixedWindow(capacity int64, rate time.Duration, fixedWindowIncrementer FixedWindowIncrementer, clock Clock) *FixedWindow {
	return &FixedWindow{backend: fixedWindowIncrementer, clock: clock, rate: rate, capacity: capacity}
}

// Limit returns the time duration to wait before the request can be processed.
// It returns ErrLimitExhausted if the request overflows the window's capacity.
func (f *FixedWindow) Limit(ctx context.Context) (time.Duration, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	now := f.clock.Now()
	window := now.Truncate(f.rate)
	if f.window != window {
		f.window = window
		f.overflow = false
	}
	ttl := f.rate - now.Sub(window)
	if f.overflow {
		// If the window is already overflowed don't increment the counter.
		return ttl, ErrLimitExhausted
	}
	c, err := f.backend.Increment(ctx, window, ttl)
	if err != nil {
		return 0, err
	}
	if c > f.capacity {
		f.overflow = true
		return ttl, ErrLimitExhausted
	}
	return 0, nil
}
