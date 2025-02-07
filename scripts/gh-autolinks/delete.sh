# GitHub CLI api
# https://cli.github.com/manual/gh_api

if [ -z "$2" ]; then
	echo "$0 <repo-name> <autolink-id>"
	exit 1
fi

gh api \
  -X DELETE \
  -H "Accept: application/vnd.github+json" \
  -H "X-GitHub-Api-Version: 2022-11-28" \
  /repos/TykTechnologies/$1/autolinks/$2
