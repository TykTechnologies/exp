# modcheck

This tool will go through the imports in go.mod and check with the
official go proxy to get a list of versions for each of the imports.

Run `modcheck` in your repo where `go.mod` exists.

Run `modcheck -suggest` for a list of dependencies that are versioned and upgradeable.

Install:

```
go install github.com/TykTechnologies/exp/cmd/modcheck@main
```

The report is provided in markdown output, suitable for github issues.

Example output for itself:

| IMPORT                 | VERSION                              | LATEST       | WARNINGS                        | CVES |
|:---|:---|:---|:---|:---|
| olekukonko/tablewriter | v0.0.6 0.20230925090304 df64c4bbad77 | v0.0.5       | Version ahead of latest release |      |
| golang.org/x/mod       | v0.14.0                              | ✓ Up to date |                                 |      |

Several warnings are printed:

- Bad request, possibly renamed (Jeffail/gabs becomes jeffail/gabs)
- Dependency without go.mod (important for plugin compiler conflicts)
- Version ahead of latest release (typically `master` or `main` branch at some commit after a tagged release)
- Deprecated import (gopkg.in is pre-go.mod and should not be used)
- CVE reports for a dependency from the [Go Vuln DB](https://vuln.go.dev/)
