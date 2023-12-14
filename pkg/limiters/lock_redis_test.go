package limiters

import (
	"context"
	"testing"
	"time"

	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedisLock(t *testing.T) {
	timeout := time.Second

	opts := &redis.UniversalOptions{
		Addrs:        []string{"localhost:6379"},
		DialTimeout:  timeout,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		PoolSize:     500,
	}

	client := redis.NewClient(opts.Simple())

	mutex := NewLockRedis(goredis.NewPool(client), "test-redsync")

	ctx := context.Background()
	assert.NoError(t, mutex.Lock(ctx))
	assert.NoError(t, mutex.Unlock(ctx))
}
