# GitHub CLI api
# https://cli.github.com/manual/gh_api

gh api \
  --method POST \
  -H "Accept: application/vnd.github+json" \
  -H "X-GitHub-Api-Version: 2022-11-28" \
  /repos/TykTechnologies/$1/autolinks \
   -f "key_prefix=TT-" -f "url_template=https://tyktech.atlassian.net/browse/<num>" -F "is_alphanumeric=true"
