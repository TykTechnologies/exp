# GitHub Actions

This folder contains developer documentation for GitHub Actions for
several repositories. This synchronizes the github workflows from:

- [tyk](./tyk)
- [tyk-analytics](./tyk-analytics)
- [tyk-analytics-ui](./tyk-analytics-ui)
- [tyk-docs](./tyk-docs)
- [exp](./exp)

To update the workflows and .md docs:

- `task clone` to clone the repositories under git/,
- `task sync` to copy workflows and run github-actions-viz.

This is not a source of truth for the workflows, as they are copied from
the individual repositories.
