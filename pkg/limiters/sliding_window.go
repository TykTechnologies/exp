package limiters

import (
	"context"
	"time"
)

// SlidingWindow implements a Sliding Window rate limiting algorithm.
//
// It does not require a distributed lock and uses a minimum amount of memory, however it will disallow all the requests
// in case when a client is flooding the service with requests.
// It's the client's responsibility to handle the disallowed request and wait before making a new request again.
type SlidingWindow struct {
	backend  SlidingWindowIncrementer
	clock    Clock
	rate     time.Duration
	capacity int64
	epsilon  float64
}

// NewSlidingWindow creates a new instance of SlidingWindow.
// Capacity is the maximum amount of requests allowed per window.
// Rate is the window size.
// Epsilon is the max-allowed range of difference when comparing the current weighted number of requests with capacity.
func NewSlidingWindow(capacity int64, rate time.Duration, slidingWindowIncrementer SlidingWindowIncrementer, clock Clock, epsilon float64) *SlidingWindow {
	return &SlidingWindow{backend: slidingWindowIncrementer, clock: clock, rate: rate, capacity: capacity, epsilon: epsilon}
}

// Limit returns the time duration to wait before the request can be processed.
// It returns ErrLimitExhausted if the request overflows the capacity.
func (s *SlidingWindow) Limit(ctx context.Context) (time.Duration, error) {
	now := s.clock.Now()
	currWindow := now.Truncate(s.rate)
	prevWindow := currWindow.Add(-s.rate)
	ttl := s.rate - now.Sub(currWindow)
	prev, curr, err := s.backend.Increment(ctx, prevWindow, currWindow, ttl+s.rate)
	if err != nil {
		return 0, err
	}

	total := float64(prev*int64(ttl))/float64(s.rate) + float64(curr)
	if total-float64(s.capacity) >= s.epsilon {
		var wait time.Duration
		if curr <= s.capacity-1 && prev > 0 {
			wait = ttl - time.Duration(float64(s.capacity-1-curr)/float64(prev)*float64(s.rate))
		} else {
			// If prev == 0.
			wait = ttl + time.Duration((1-float64(s.capacity-1)/float64(curr))*float64(s.rate))
		}
		return wait, ErrLimitExhausted
	}
	return 0, nil
}
