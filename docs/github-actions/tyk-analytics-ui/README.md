# Using GitHub Actions

This directory contains the workflows that are executed for the purposes of CI.

## test-react.yml

Run a matrix test on ubuntu-16.04, ubuntu-latest and node-10.x and node-12.x. The PR _must_ pass the 12.x test on ubuntu-latest.

The pass/fail criteria are in "Settings->Branches->Branch protection rules" (needs admin).

The private npm modules are fetched using the auth token used in the build pipelines. `NPM_TOKEN` is a Github secret on this repo.

## update-tyk-analytics.yml

Fire a `build-assets` event to the [repository-dispatch](https://help.github.com/en/actions/reference/events-that-trigger-workflows#external-events-repository_dispatch) pseudo-webhook.

This uses a repo scoped personal access token and is saved as a Github secret in this repo.

## Documentation

[Workflow reference](https://help.github.com/en/actions/reference/workflow-syntax-for-github-actions)
[Triggers](https://help.github.com/en/actions/reference/events-that-trigger-workflows) for workflows

