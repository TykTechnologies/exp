# Contents

- [Detect memory leak exposure](./memory-leaks)
- [Detect non-returning loops](./non-returning-loop)
- [Detect symbol usage code smells](./usage)
- [Repeated Logger calls in scope](./repeated-logger)

# Structure

Structure rule or rulesets under folders below.

- Semgrep rules: `rules.yml` for each rule,
- Taskfile to run scans: `Taskfile.yml`,
- Results of the scan: `output.json` (indented),
- Scan example code: `example.go`.

Run `task` in each folder to run.
