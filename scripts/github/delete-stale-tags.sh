#!/bin/bash

OWNER=${1:-TykTechnologies}
REPO=${2:-tyk}
PATTERN=".*-.*alpha.*"

# List and filter tags
TAGS=$(gh api repos/$OWNER/$REPO/tags --paginate | jq -r --arg pattern "$PATTERN" '.[] | select(.name | test($pattern)) | .name')

for TAG in $TAGS; do
  echo "# delete tag $TAG" >&2
  echo gh api -X DELETE repos/$OWNER/$REPO/git/refs/tags/$TAG
done
