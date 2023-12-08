package limiters

import (
	"context"
	"sync"
	"time"
)

// SlidingWindowInMemory is an in-memory implementation of SlidingWindowIncrementer.
type SlidingWindowInMemory struct {
	mu           sync.Mutex
	prevC, currC int64
	prevW, currW time.Time
}

// NewSlidingWindowInMemory creates a new instance of SlidingWindowInMemory.
func NewSlidingWindowInMemory() *SlidingWindowInMemory {
	return &SlidingWindowInMemory{}
}

// Increment increments the current window's counter and returns the number of requests in the previous window and the
// current one.
func (s *SlidingWindowInMemory) Increment(ctx context.Context, prev, curr time.Time, _ time.Duration) (int64, int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if curr != s.currW {
		if prev == s.currW {
			s.prevW = s.currW
			s.prevC = s.currC
		} else {
			s.prevW = time.Time{}
			s.prevC = 0
		}
		s.currW = curr
		s.currC = 0
	}
	s.currC++
	return s.prevC, s.currC, ctx.Err()
}
