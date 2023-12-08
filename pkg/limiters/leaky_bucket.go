package limiters

import (
	"context"
	"sync"
	"time"
)

// LeakyBucket implements the https://en.wikipedia.org/wiki/Leaky_bucket#As_a_queue algorithm.
type LeakyBucket struct {
	locker  DistLocker
	backend LeakyBucketStateBackend
	clock   Clock
	logger  Logger
	// Capacity is the maximum allowed number of tockens in the bucket.
	capacity int64
	// Rate is the output rate: 1 request per the rate duration (in nanoseconds).
	rate int64
	mu   sync.Mutex
}

// NewLeakyBucket creates a new instance of LeakyBucket.
func NewLeakyBucket(capacity int64, rate time.Duration, locker DistLocker, leakyBucketStateBackend LeakyBucketStateBackend, clock Clock, logger Logger) *LeakyBucket {
	return &LeakyBucket{
		locker:   locker,
		backend:  leakyBucketStateBackend,
		clock:    clock,
		logger:   logger,
		capacity: capacity,
		rate:     int64(rate),
	}
}

// Limit returns the time duration to wait before the request can be processed.
// If the last request happened earlier than the rate this method returns zero duration.
// It returns ErrLimitExhausted if the the request overflows the bucket's capacity. In this case the returned duration
// means how long it would have taken to wait for the request to be processed if the bucket was not overflowed.
func (t *LeakyBucket) Limit(ctx context.Context) (time.Duration, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if err := t.locker.Lock(ctx); err != nil {
		return 0, err
	}
	defer func() {
		if err := t.locker.Unlock(ctx); err != nil {
			t.logger.Log(err)
		}
	}()
	state, err := t.backend.State(ctx)
	if err != nil {
		return 0, err
	}
	now := t.clock.Now().UnixNano()
	if now < state.Last {
		// The queue has requests in it: move the current request to the last position + 1.
		state.Last += t.rate
	} else {
		// The queue is empty.
		// The offset is the duration to wait in case the last request happened less than rate duration ago.
		var offset int64
		delta := now - state.Last
		if delta < t.rate {
			offset = t.rate - delta
		}
		state.Last = now + offset
	}

	wait := state.Last - now
	if wait/t.rate > t.capacity {
		return time.Duration(wait), ErrLimitExhausted
	}
	if err = t.backend.SetState(ctx, state); err != nil {
		return 0, err
	}
	return time.Duration(wait), nil
}
