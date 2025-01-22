declare -A details

details['BenchmarkWrappedServeHTTP']=true

declare -A releases

# Releases is an associative array between the release tag,
# and a release branch. The key is usually the latest patch
# release that was tagged, but can even be a RC tag.

releases['v5.0.15']="release-5-lts"
releases['v5.1.2']="release-5.1"
releases['v5.2.6']="release-5.2"
releases['v5.3.9']="release-5.3"
