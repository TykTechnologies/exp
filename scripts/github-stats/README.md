# Github actions stats

1. Adjust for your repo and run `./get-actions-runs.sh`
2. Run `go run .` to process the data/ file

The json is processed as such:

- filters for `.github/workflows/ci-test.yml`
- skips workflows that ran longer than 1h
- aggregates users by name/email
- does a bit of math to calculate time spend/time saved ROI

Adjust the script/main.go file to your repository needs.