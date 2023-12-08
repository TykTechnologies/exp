package limiters

// LeakyBucketState represents the state of a LeakyBucket.
type LeakyBucketState struct {
	// Last is the Unix timestamp in nanoseconds of the most recent request.
	Last int64
}

// IzZero returns true if the bucket state is zero valued.
func (s LeakyBucketState) IzZero() bool {
	return s.Last == 0
}
