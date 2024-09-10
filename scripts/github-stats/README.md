# Github actions stats

1. Adjust for your repo and run `./get-actions-runs.sh`
2. Run `task` to process the data/ file

The `cmd/monthly` processes the json so:

- filters for `.github/workflows/ci-test.yml`
- filter 30 days
- group by day and branch
- exclude 0 failure branches

The `cmd/authors` processes the json so:

- filters for `.github/workflows/ci-test.yml`
- skips workflows that ran longer than 1h
- aggregates users by name/email
- does a bit of math to calculate time spend/time saved ROI
