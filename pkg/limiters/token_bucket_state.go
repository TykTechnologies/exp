package limiters

// TokenBucketState represents a state of a token bucket.
type TokenBucketState struct {
	// Last is the last time the state was updated (Unix timestamp in nanoseconds).
	Last int64
	// Available is the number of available tokens in the bucket.
	Available int64
}

// isZero returns true if the bucket state is zero valued.
func (s TokenBucketState) isZero() bool {
	return s.Last == 0 && s.Available == 0
}
