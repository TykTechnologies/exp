package limiters

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// LeakyBucketRedis is a Redis implementation of a LeakyBucketStateBackend.
type LeakyBucketRedis struct {
	cli         redis.UniversalClient
	prefix      string
	ttl         time.Duration
	raceCheck   bool
	lastVersion int64
}

// NewLeakyBucketRedis creates a new LeakyBucketRedis instance.
// Prefix is the key prefix used to store all the keys used in this implementation in Redis.
// TTL is the TTL of the stored keys.
//
// If raceCheck is true and the keys in Redis are modified in between State() and SetState() calls then
// ErrRaceCondition is returned.
func NewLeakyBucketRedis(cli redis.UniversalClient, prefix string, ttl time.Duration, raceCheck bool) *LeakyBucketRedis {
	return &LeakyBucketRedis{cli: cli, prefix: prefix, ttl: ttl, raceCheck: raceCheck}
}

// SetState updates the state in Redis.
// The provided fencing token is checked on the Redis side before saving the keys.
func (t *LeakyBucketRedis) SetState(ctx context.Context, state LeakyBucketState) error {
	var err error
	done := make(chan struct{}, 1)
	go func() {
		defer close(done)
		if !t.raceCheck {
			err = t.cli.Set(ctx, redisKey(t.prefix, redisKeyLBLast), state.Last, t.ttl).Err()
			return
		}
		var result interface{}
		// TODO: make use of EVALSHA.
		result, err = t.cli.Eval(ctx, `
	local version = tonumber(redis.call('get', KEYS[1])) or 0
	if version > tonumber(ARGV[1]) then
		return 'RACE_CONDITION'
	end
	return {
		redis.call('incr', KEYS[1]),
		redis.call('pexpire', KEYS[1], ARGV[3]),
		redis.call('set', KEYS[2], ARGV[2], 'PX', ARGV[3]),
	}
	`, []string{
			redisKey(t.prefix, redisKeyLBVersion),
			redisKey(t.prefix, redisKeyLBLast),
		},
			t.lastVersion,
			state.Last,
			// TTL in milliseconds.
			int64(t.ttl/time.Microsecond)).Result()

		if err == nil {
			err = checkResponseFromRedis(result, []interface{}{t.lastVersion + 1, int64(1), "OK"})
		}
	}()

	select {
	case <-done:

	case <-ctx.Done():
		return ctx.Err()
	}

	return errorsWrap(err, "failed to save keys to redis")
}

// State gets the bucket's state from Redis.
func (t *LeakyBucketRedis) State(ctx context.Context) (LeakyBucketState, error) {
	var values []interface{}
	var err error
	done := make(chan struct{}, 1)
	go func() {
		defer close(done)
		keys := []string{
			redisKey(t.prefix, redisKeyLBLast),
		}
		if t.raceCheck {
			keys = append(keys, redisKey(t.prefix, redisKeyLBVersion))
		}
		values, err = t.cli.MGet(ctx, keys...).Result()
	}()

	select {
	case <-done:

	case <-ctx.Done():
		return LeakyBucketState{}, ctx.Err()
	}

	if err != nil {
		return LeakyBucketState{}, errorsWrap(err, "failed to get keys from redis")
	}
	nilAny := false
	for _, v := range values {
		if v == nil {
			nilAny = true
			break
		}
	}
	if nilAny || err == redis.Nil {
		// Keys don't exist, return an empty state.
		return LeakyBucketState{}, nil
	}

	last, err := strconv.ParseInt(values[0].(string), 10, 64)
	if err != nil {
		return LeakyBucketState{}, err
	}
	if t.raceCheck {
		t.lastVersion, err = strconv.ParseInt(values[1].(string), 10, 64)
		if err != nil {
			return LeakyBucketState{}, err
		}
	}
	return LeakyBucketState{
		Last: last,
	}, nil
}
