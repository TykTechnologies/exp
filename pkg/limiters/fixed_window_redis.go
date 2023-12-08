package limiters

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// FixedWindowRedis implements FixedWindow in Redis.
type FixedWindowRedis struct {
	cli    *redis.Client
	prefix string
}

// NewFixedWindowRedis returns a new instance of FixedWindowRedis.
// Prefix is the key prefix used to store all the keys used in this implementation in Redis.
func NewFixedWindowRedis(cli *redis.Client, prefix string) *FixedWindowRedis {
	return &FixedWindowRedis{cli: cli, prefix: prefix}
}

// Increment increments the window's counter in Redis.
func (f *FixedWindowRedis) Increment(ctx context.Context, window time.Time, ttl time.Duration) (int64, error) {
	var incr *redis.IntCmd
	var err error
	done := make(chan struct{})
	go func() {
		defer close(done)
		_, err = f.cli.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
			key := fmt.Sprintf("%d", window.UnixNano())
			incr = pipeliner.Incr(ctx, redisKey(f.prefix, key))
			pipeliner.PExpire(ctx, redisKey(f.prefix, key), ttl)
			return nil
		})
	}()

	select {
	case <-done:
		if err != nil {
			return 0, errorsWrap(err, "redis transaction failed")
		}
		return incr.Val(), incr.Err()
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}
