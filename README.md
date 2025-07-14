# Experimental and development tooling

This repository mimics a subset of
[golang.org/x/exp](https://pkg.go.dev/golang.org/x/exp), namely:

- `/cmd` - various development tooling and experiments

In particular tools like `go-fsck`, `schema-gen` are usable for code
analysis and linting with an extended godoc ruleset applying to fields.
This is used to lint our data model in several places. The `summary` tool
summarizes various things and also has some adoption.

## Aim of the repository

The aim of the repository is to provide a space for the development of tooling
that aids common CI and data flows at Tyk. Common tools that fall in this area are:

- Code generation tooling,
- Code inspection / analysis to surface schema,
- Converters between schema formats,
- Generates documentation from schema,
- Static code analysis,
- Go Reflection and AST tooling,
- Refactoring tooling,
- Large scale change tooling,
- Toy example projects,
- Full scale test suites,
- Linters, etc...

Individual projects must be self-contained within `/cmd`, and may provide
their own go.mod to decrease scope if a shared dependency tree becomes a
problem. Each tool needs to pass a build and a test check. There isn't a
requirement for code coverage, but this will eventually become the only
planned merge check.

Common files that a command is encouraged to provide:

1. README.md
    - explain why this exists,
    - who is it for,
    - document stability,
    - list usage if it's enabled in CI, link github actions,
    - list ownership / point of contact,
    - document ANY users that want to be contacted if a breaking change happens
2. Taskfile.yml or Makefile? - collection of targets for the project, required: `build`, `test`.
    - We can auto default to `go build` and `go test -race -count=100 ./...`.
      This seems strict, but we have 0 legacy to care. And you really have to
      think which test you're going to write. Code coverage is NOT a requirement,
      however there are exceptions where we pay attention. See step 1.
3. `main.go`, `internal` package, `go.mod` (no metaversion, we're v0 all the way)
4. `Dockerfile`, `docker-compose.yml`

Take a look around and figure out which of these apply to you.

Any added github workflows are expected to run against the last two Go
versions released. At the time of writing, this would mean 1.19 and 1.20.

## Issues

You own a cmd/, you maintain the cmd to your own standards. You don't
need to consider github issues or accept PRs, but this is an open source
shop, sir. Make it clear in your README what kind of issues and contact
you are willing to accept, if any. Nothing wrong with saying no.

## Compatibility promises

Warning: Packages here are experimental and unreliable. Some may one day
be promoted to the main repository or other subrepository, or they may be
modified arbitrarily or even disappear altogether.

## GitHub Actions Workflows

This repository provides several GitHub Actions workflows to automate common tasks across Tyk repositories.

### Code Freeze Branch Creation

The `code-freeze.yml` workflow creates consistent branches across multiple Tyk repositories simultaneously.

**Purpose**: Automate the creation of release branches across the Tyk ecosystem.

**Repositories affected**:
- tyk-analytics
- tyk-analytics-ui
- tyk

**Usage**:
1. Go to the Actions tab in the repository
2. Select "Code Freeze Branch Creation" workflow
3. Click "Run workflow"
4. Fill in the parameters:
   - Source Branch: The branch to create from (e.g., master)
   - Destination Branch: The branch to create (e.g., release-5.8)
   - Send Slack notification: Whether to notify the team via Slack

**Example use case**: Creating a release-5.8 branch from master across all repositories at once.

### Create Tags Across Repositories

The `create-tags.yml` workflow creates consistent tags across multiple Tyk repositories simultaneously with flexible repository selection.

**Purpose**: Automate the creation of release tags across selected repositories in the Tyk ecosystem.

**Available repositories**:
- tyk
- tyk-analytics
- tyk-analytics-ui
- tyk-sink
- tyk-sync-internal
- tyk-operator-internal
- tyk-charts
- portal
- tyk-pump

**Usage**:
1. Go to the Actions tab in the repository
2. Select "Create Tags Across Repositories" workflow
3. Click "Run workflow"
4. Fill in the parameters:
   - Source Branch: The branch to create the tag from (e.g., master, release-5.8)
   - Tag Name: The name of the tag to create (e.g., v5.8.0)
   - Tag Message: Optional message for annotated tags
   - Send Slack notification: Whether to notify the team via Slack
   - Repository checkboxes: Select which repositories to create tags for

**Key features**:
- **Selective repository targeting**: Choose specific repositories via checkboxes instead of tagging all repositories
- **Validation**: Ensures at least one repository is selected before proceeding
- **GitHub API integration**: Uses GitHub API for reliable tag creation
- **Enhanced error handling**: Provides clear feedback for each repository operation

**Example use case**: Creating a v5.8.0 tag from the release-5.8 branch on only tyk, tyk-analytics, and tyk-pump repositories.

**Note**: Both workflows use the `ORG_GH_TOKEN` secret for authentication and the `UI_SLACK_AUTH_TOKEN` secret for Slack notifications. Ensure these secrets are properly configured in the repository settings.