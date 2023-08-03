# GitHub Actions

## build-assets.yml

This is fired from [tyk-analytics-ui](https://github.com/TykTechnologies/tyk-analytics-ui).

- Fetches the specified commit of webclient
- Builds webclient
- Updates go-bindata in host repo tyk-analytics
- Commit webclient and bindata
- Push changes

Uses a deploy key. The priv key is a Github secret called `bindata_deploy_key`.

Uses setup-go@v2-beta so that PATH is setup for go-bindata.

## ci-tests.yml

Runs on all PRs against master. This is a required workflow for merging.

## integration-image.yml

Runs on all pushes to `integration/*`. 

See https://github.com/TykTechnologies/tyk-analytics/wiki/Integration-Testing

