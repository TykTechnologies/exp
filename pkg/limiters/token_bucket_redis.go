package limiters

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// TokenBucketRedis is a Redis implementation of a TokenBucketStateBackend.
//
// Redis is an in-memory key-value data storage which also supports persistence.
// It is a better choice for high load cases than etcd as it does not keep old values of the keys thus does not need
// the compaction/defragmentation.
//
// Although depending on a persistence and a cluster configuration some data might be lost in case of a failure
// resulting in an under-limiting the accesses to the service.
type TokenBucketRedis struct {
	cli         redis.UniversalClient
	prefix      string
	ttl         time.Duration
	raceCheck   bool
	lastVersion int64
}

// NewTokenBucketRedis creates a new TokenBucketRedis instance.
// Prefix is the key prefix used to store all the keys used in this implementation in Redis.
// TTL is the TTL of the stored keys.
//
// If raceCheck is true and the keys in Redis are modified in between State() and SetState() calls then
// ErrRaceCondition is returned.
// This adds an extra overhead since a Lua script has to be executed on the Redis side which locks the entire database.
func NewTokenBucketRedis(cli redis.UniversalClient, prefix string, ttl time.Duration, raceCheck bool) *TokenBucketRedis {
	return &TokenBucketRedis{cli: cli, prefix: prefix, ttl: ttl, raceCheck: raceCheck}
}

// SetState updates the state in Redis.
func (t *TokenBucketRedis) SetState(ctx context.Context, state TokenBucketState) error {
	var err error
	done := make(chan struct{}, 1)
	go func() {
		defer close(done)
		if !t.raceCheck {
			_, err = t.cli.TxPipelined(ctx, func(pipeliner redis.Pipeliner) error {
				if err = pipeliner.Set(ctx, redisKey(t.prefix, redisKeyTBLast), state.Last, t.ttl).Err(); err != nil {
					return err
				}
				return pipeliner.Set(ctx, redisKey(t.prefix, redisKeyTBAvailable), state.Available, t.ttl).Err()
			})
			return
		}
		var result interface{}
		// TODO: use EVALSHA.
		result, err = t.cli.Eval(ctx, `
	local version = tonumber(redis.call('get', KEYS[1])) or 0
	if version > tonumber(ARGV[1]) then
		return 'RACE_CONDITION'
	end
	return {
		redis.call('incr', KEYS[1]),
		redis.call('pexpire', KEYS[1], ARGV[4]),
		redis.call('set', KEYS[2], ARGV[2], 'PX', ARGV[4]),
		redis.call('set', KEYS[3], ARGV[3], 'PX', ARGV[4]),
	}
	`, []string{
			redisKey(t.prefix, redisKeyTBVersion),
			redisKey(t.prefix, redisKeyTBLast),
			redisKey(t.prefix, redisKeyTBAvailable),
		},
			t.lastVersion,
			state.Last,
			state.Available,
			// TTL in milliseconds.
			int64(t.ttl/time.Millisecond)).Result()

		if err == nil {
			err = checkResponseFromRedis(result, []interface{}{t.lastVersion + 1, int64(1), "OK", "OK"})
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
func (t *TokenBucketRedis) State(ctx context.Context) (TokenBucketState, error) {
	var values []interface{}
	var err error
	done := make(chan struct{}, 1)

	if t.raceCheck {
		// reset in a case of returning an empty TokenBucketState
		t.lastVersion = 0
	}

	go func() {
		defer close(done)
		keys := []string{
			redisKey(t.prefix, redisKeyTBLast),
			redisKey(t.prefix, redisKeyTBAvailable),
		}
		if t.raceCheck {
			keys = append(keys, redisKey(t.prefix, redisKeyTBVersion))
		}
		values, err = t.cli.MGet(ctx, keys...).Result()
	}()

	select {
	case <-done:

	case <-ctx.Done():
		return TokenBucketState{}, ctx.Err()
	}

	if err != nil {
		return TokenBucketState{}, errorsWrap(err, "failed to get keys from redis")
	}
	nilAny := false
	for _, v := range values {
		if v == nil {
			nilAny = true
			break
		}
	}
	if nilAny || err == redis.Nil {
		// Keys don't exist, return the initial state.
		return TokenBucketState{}, nil
	}

	last, err := strconv.ParseInt(values[0].(string), 10, 64)
	if err != nil {
		return TokenBucketState{}, err
	}
	available, err := strconv.ParseInt(values[1].(string), 10, 64)
	if err != nil {
		return TokenBucketState{}, err
	}
	if t.raceCheck {
		t.lastVersion, err = strconv.ParseInt(values[2].(string), 10, 64)
		if err != nil {
			return TokenBucketState{}, err
		}
	}
	return TokenBucketState{
		Last:      last,
		Available: available,
	}, nil
}
