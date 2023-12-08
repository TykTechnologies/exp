package limiters

import (
	"context"
)

// LeakyBucketInMemory is an in-memory implementation of LeakyBucketStateBackend.
type LeakyBucketInMemory struct {
	state LeakyBucketState
}

// NewLeakyBucketInMemory creates a new instance of LeakyBucketInMemory.
func NewLeakyBucketInMemory() *LeakyBucketInMemory {
	return &LeakyBucketInMemory{}
}

// SetState sets the current state of the bucket.
func (l *LeakyBucketInMemory) SetState(ctx context.Context, state LeakyBucketState) error {
	l.state = state
	return ctx.Err()
}

// State gets the current state of the bucket.
func (l *LeakyBucketInMemory) State(ctx context.Context) (LeakyBucketState, error) {
	return l.state, ctx.Err()
}
