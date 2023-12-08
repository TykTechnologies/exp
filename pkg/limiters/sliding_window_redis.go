package limiters

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// SlidingWindowRedis implements SlidingWindow in Redis.
type SlidingWindowRedis struct {
	cli    *redis.Client
	prefix string
}

// NewSlidingWindowRedis creates a new instance of SlidingWindowRedis.
func NewSlidingWindowRedis(cli *redis.Client, prefix string) *SlidingWindowRedis {
	return &SlidingWindowRedis{cli: cli, prefix: prefix}
}

// Increment increments the current window's counter in Redis and returns the number of requests in the previous window
// and the current one.
func (s *SlidingWindowRedis) Increment(ctx context.Context, prev, curr time.Time, ttl time.Duration) (int64, int64, error) {
	var incr *redis.IntCmd
	var prevCountCmd *redis.StringCmd
	var err error
	done := make(chan struct{})
	go func() {
		defer close(done)
		_, err = s.cli.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
			currKey := fmt.Sprintf("%d", curr.UnixNano())
			incr = pipeliner.Incr(ctx, redisKey(s.prefix, currKey))
			pipeliner.PExpire(ctx, redisKey(s.prefix, currKey), ttl)
			prevCountCmd = pipeliner.Get(ctx, redisKey(s.prefix, fmt.Sprintf("%d", prev.UnixNano())))
			return nil
		})
	}()

	var prevCount int64
	select {
	case <-done:
		if err == redis.TxFailedErr {
			return 0, 0, errorsWrap(err, "redis transaction failed")
		} else if err == redis.Nil {
			prevCount = 0
		} else if err != nil {
			return 0, 0, errorsWrap(err, "unexpected error from redis")
		} else {
			prevCount, err = strconv.ParseInt(prevCountCmd.Val(), 10, 64)
			if err != nil {
				return 0, 0, errorsWrap(err, "failed to parse response from redis")
			}
		}
		return prevCount, incr.Val(), nil
	case <-ctx.Done():
		return 0, 0, ctx.Err()
	}
}
