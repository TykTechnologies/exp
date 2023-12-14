# Limiters

This is a fork of [mennanov/limiters](https://github.com/mennanov/limiters).

It has been created using go-fsck extract/restore, to group 1 struct per file.

Changes:

- Removed Concurrent Buffer and related registry,
- Removed all etcd, consul, dynamodb implementations,
- Added an errorsWrap function to use the standar lib errors pkg,
- Updated code to use stdlib errors over github.com/pkg/errors,
- Removed SystemClock.Sleep,
- Added per-key local storage constructors
- Used `redis.UniversalClient` instead of `*redis.Client` to support cluster
- Adjusted RedisLock to use Lock/UnlockContext functions from redsync
- Wrote test for RedisLock

Ultimately the fork contains redis and local implementations for:

- Token bucket
- Leaky bucket
- Fixed window counter
- Sliding window counter

It only imports the redis and redsync dependencies. These are used to
implement the distributed rate limiters, and provide a distributed
locking mechanism for the rate limiters that require it.
