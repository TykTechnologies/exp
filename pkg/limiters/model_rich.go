package limiters

import (
	"context"
	"time"
)

// ConcurrentBufferBackend wraps the Add and Remove methods.
type ConcurrentBufferBackend interface {
	// Add adds the request with the given key to the buffer and returns the total number of requests in it.
	Add(ctx context.Context, key string) (int64, error)
	// Remove removes the request from the buffer.
	Remove(ctx context.Context, key string) error
}

// DistLocker is a context aware distributed locker (interface is similar to sync.Locker).
type DistLocker interface {
	// Lock locks the locker.
	Lock(ctx context.Context) error
	// Unlock unlocks the previously successfully locked lock.
	Unlock(ctx context.Context) error
}

// FixedWindowIncrementer wraps the Increment method.
type FixedWindowIncrementer interface {
	// Increment increments the request counter for the window and returns the counter value.
	// TTL is the time duration before the next window.
	Increment(ctx context.Context, window time.Time, ttl time.Duration) (int64, error)
}

// LeakyBucketStateBackend interface encapsulates the logic of retrieving and persisting the state of a LeakyBucket.
type LeakyBucketStateBackend interface {
	// State gets the current state of the LeakyBucket.
	State(ctx context.Context) (LeakyBucketState, error)
	// SetState sets (persists) the current state of the LeakyBucket.
	SetState(ctx context.Context, state LeakyBucketState) error
}

// SlidingWindowIncrementer wraps the Increment method.
type SlidingWindowIncrementer interface {
	// Increment increments the request counter for the current window and returns the counter values for the previous
	// window and the current one.
	// TTL is the time duration before the next window.
	Increment(ctx context.Context, prev, curr time.Time, ttl time.Duration) (prevCount, currCount int64, err error)
}

// TokenBucketStateBackend interface encapsulates the logic of retrieving and persisting the state of a TokenBucket.
type TokenBucketStateBackend interface {
	// State gets the current state of the TokenBucket.
	State(ctx context.Context) (TokenBucketState, error)
	// SetState sets (persists) the current state of the TokenBucket.
	SetState(ctx context.Context, state TokenBucketState) error
}
