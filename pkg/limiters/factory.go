package limiters

import (
	"sync"
)

// This file contains the factory methods that allocate a key based
// in memory storage for the various rate algos memory storages.

var globalMu sync.Mutex
var global struct {
	FixedWindow   map[string]*FixedWindowInMemory
	SlidingWindow map[string]*SlidingWindowInMemory
	LeakyBucket   map[string]*LeakyBucketInMemory
	TokenBucket   map[string]*TokenBucketInMemory
}

func LocalFixedWindow(key string) *FixedWindowInMemory {
	globalMu.Lock()
	defer globalMu.Unlock()

	ctor := NewFixedWindowInMemory

	if global.FixedWindow == nil {
		global.FixedWindow = map[string]*FixedWindowInMemory{
			key: ctor(),
		}
	}

	val, ok := global.FixedWindow[key]
	if !ok {
		val = ctor()
		global.FixedWindow[key] = val
	}
	return val
}

func LocalSlidingWindow(key string) *SlidingWindowInMemory {
	globalMu.Lock()
	defer globalMu.Unlock()

	ctor := NewSlidingWindowInMemory

	if global.SlidingWindow == nil {
		global.SlidingWindow = map[string]*SlidingWindowInMemory{
			key: ctor(),
		}
	}

	val, ok := global.SlidingWindow[key]
	if !ok {
		val = ctor()
		global.SlidingWindow[key] = val
	}
	return val
}

func LocalLeakyBucket(key string) *LeakyBucketInMemory {
	globalMu.Lock()
	defer globalMu.Unlock()

	ctor := NewLeakyBucketInMemory

	if global.LeakyBucket == nil {
		global.LeakyBucket = map[string]*LeakyBucketInMemory{
			key: ctor(),
		}
	}

	val, ok := global.LeakyBucket[key]
	if !ok {
		val = ctor()
		global.LeakyBucket[key] = val
	}
	return val
}

func LocalTokenBucket(key string) *TokenBucketInMemory {
	globalMu.Lock()
	defer globalMu.Unlock()

	ctor := NewTokenBucketInMemory

	if global.TokenBucket == nil {
		global.TokenBucket = map[string]*TokenBucketInMemory{
			key: ctor(),
		}
	}

	val, ok := global.TokenBucket[key]
	if !ok {
		val = ctor()
		global.TokenBucket[key] = val
	}
	return val
}
