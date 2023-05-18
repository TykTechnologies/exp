# Experimental and development tooling

This repository mimics a subset of
[golang.org/x/exp](https://pkg.go.dev/golang.org/x/exp), namely:

- `/cmd` - various development tooling and experiments

The focus area for this repository is to provide a scratch space for
tooling aimed at supporting development at Tyk. Individual projects must
be self-contained within `/cmd`, and may provide their own go.mod to
decrease scope if a shared dependency tree becomes a problem.

In general, the tooling, regardless of the experimental nature of the
development, needs to pass a build and a test check. There isn't a
requirement for code coverage, but this will eventually become the only
planned merge check. Each cmd package will be tested individually.

Common files that a command is encouraged to provide:

1. README.md
    - explain why this exists,
    - who is it for,
    - document stability, if it's enabled in CI, link github actions,
    - list ownership / point of contact,
    - document ANY users that want to be contacted if a breaking change happens
2. Taskfile.yml - a collection of commands for the project, required: `build`, `test`.
    - We can auto default to `go build` and `go test -race -count=100 ./...`.
      This seems strict, but we have 0 legacy to care. And you really have to
      think which test you're going to write. Code coverage is NOT a requirement,
      however there are exceptions where we pay attention. See step 1.
3. `main.go`, `internal` package, `go.mod` (no metaversion, we're v0 all the way)
4. `Dockerfile`, `docker-compose.yml`

I'd say look around and figure out which of these apply to you.

Any added github workflows are expected to run against the last two Go
versions released. At the time of writing, this would mean 1.19 and 1.20.

## Compatibility promises

Warning: Packages here are experimental and unreliable. Some may one day
be promoted to the main repository or other subrepository, or they may be
modified arbitrarily or even disappear altogether.
