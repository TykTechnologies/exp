package limiters

import (
	"context"
)

// TokenBucketInMemory is an in-memory implementation of TokenBucketStateBackend.
//
// The state is not shared nor persisted so it won't survive restarts or failures.
// Due to the local nature of the state the rate at which some endpoints are accessed can't be reliably predicted or
// limited.
//
// Although it can be used as a global rate limiter with a round-robin load-balancer.
type TokenBucketInMemory struct {
	state TokenBucketState
}

// NewTokenBucketInMemory creates a new instance of TokenBucketInMemory.
func NewTokenBucketInMemory() *TokenBucketInMemory {
	return &TokenBucketInMemory{}
}

// SetState sets the current bucket's state.
func (t *TokenBucketInMemory) SetState(ctx context.Context, state TokenBucketState) error {
	t.state = state
	return ctx.Err()
}

// State returns the current bucket's state.
func (t *TokenBucketInMemory) State(ctx context.Context) (TokenBucketState, error) {
	return t.state, ctx.Err()
}
