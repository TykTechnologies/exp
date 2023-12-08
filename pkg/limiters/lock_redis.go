package limiters

import (
	"context"

	"github.com/go-redsync/redsync/v4"
	redsyncredis "github.com/go-redsync/redsync/v4/redis"
)

// LockRedis is a wrapper around github.com/go-redsync/redsync that implements the DistLocker interface.
type LockRedis struct {
	mutex *redsync.Mutex
}

// NewLockRedis creates a new instance of LockRedis.
func NewLockRedis(pool redsyncredis.Pool, mutexName string) *LockRedis {
	rs := redsync.New(pool)
	mutex := rs.NewMutex(mutexName)
	return &LockRedis{mutex: mutex}
}

// Lock locks the lock in Redis.
func (l *LockRedis) Lock(_ context.Context) error {
	err := l.mutex.Lock()
	return errorsWrap(err, "failed to lock a mutex in redis")
}

// Unlock unlocks the lock in Redis.
func (l *LockRedis) Unlock(_ context.Context) error {
	if ok, err := l.mutex.Unlock(); !ok || err != nil {
		return errorsWrap(err, "failed to unlock a mutex in redis")
	}
	return nil
}
