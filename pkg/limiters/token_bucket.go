package limiters

import (
	"context"
	"sync"
	"time"
)

// TokenBucket implements the https://en.wikipedia.org/wiki/Token_bucket algorithm.
type TokenBucket struct {
	locker  DistLocker
	backend TokenBucketStateBackend
	clock   Clock
	logger  Logger
	// refillRate is the tokens refill rate (1 token per duration).
	refillRate time.Duration
	// capacity is the bucket's capacity.
	capacity int64
	mu       sync.Mutex
}

// NewTokenBucket creates a new instance of TokenBucket.
func NewTokenBucket(capacity int64, refillRate time.Duration, locker DistLocker, tokenBucketStateBackend TokenBucketStateBackend, clock Clock, logger Logger) *TokenBucket {
	return &TokenBucket{
		locker:     locker,
		backend:    tokenBucketStateBackend,
		clock:      clock,
		logger:     logger,
		refillRate: refillRate,
		capacity:   capacity,
	}
}

// Limit takes 1 token from the bucket.
func (t *TokenBucket) Limit(ctx context.Context) (time.Duration, error) {
	return t.Take(ctx, 1)
}

// Take takes tokens from the bucket.
//
// It returns a zero duration and a nil error if the bucket has sufficient amount of tokens.
//
// It returns ErrLimitExhausted if the amount of available tokens is less than requested. In this case the returned
// duration is the amount of time to wait to retry the request.
func (t *TokenBucket) Take(ctx context.Context, tokens int64) (time.Duration, error) {
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
	if state.isZero() {
		// Initially the bucket is full.
		state.Available = t.capacity
	}
	now := t.clock.Now().UnixNano()
	// Refill the bucket.
	tokensToAdd := (now - state.Last) / int64(t.refillRate)
	partialTime := (now - state.Last) % int64(t.refillRate)
	if tokensToAdd > 0 {
		if tokensToAdd+state.Available < t.capacity {
			state.Available += tokensToAdd
			state.Last = now - partialTime
		} else {
			state.Available = t.capacity
			state.Last = now
		}
	}

	if tokens > state.Available {
		return t.refillRate * time.Duration(tokens-state.Available), ErrLimitExhausted
	}
	// Take the tokens from the bucket.
	state.Available -= tokens
	if err = t.backend.SetState(ctx, state); err != nil {
		return 0, err
	}
	return 0, nil
}
