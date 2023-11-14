#!/bin/sh
mkdir -p data
gh api -X GET /repos/TykTechnologies/tyk/actions/runs --paginate > data/github-actions-runs.json
