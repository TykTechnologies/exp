# Benchmarking and verifying rate limit behaviour

Currently this does the following:

- Start a gateway with a rate limiter configuration flag via env.
- Issue hey at a request rate of 50 requests/s for 10 seconds.
- Log the incoming request rate and HTTP status code responses with rakyll/hey.
- Log the back-end request rate and HTTP responses (200 OK, can assert JSON).
- Log the gateway `--memoryprofile` during the benchmark.
- Log the gateway `--cpuprofile` during the benchmark.

From these `.json` and `.csv` files, a report should be generated.

The API declares a rate=40, per=1.

The state of everything is pretty much work in progress, but:

- Run `task` to run all the rate limiters (needs internal gw build from POC branch to exist),
- Go to `cmd/parse` and run `task` to sanitize csv/json benchmarks to combined format,
- Go to `cmd/render` and run `task` to render combined format files to pngs.

For each benchmark the incoming rate and outgoing rate is logged. The
request/sec is calculated with a sliding log strategy considering the
last second of requests. Given that requests are blocked by the rate
limiters, the intervals where the back-end is not getting requests are
not visible (e.g. fixed window). This needs a slightly different
rendering.

Currently doesn't cover DRL or RRL (no-sentinel). RRL behaviour should
match sentinel for the benchmark, the difference would be if the input
rate dropped under the limited value. DRL should tentatively match
leaky bucket.
