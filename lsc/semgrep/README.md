# Large scale code changes with semgrep

task: Available tasks for this project:

* default:       Run semgrep
* pull:          Pull latest returntocorp/semgrep
* scan:          Scan with upstream rules

In order to use this, clone the project you're working on into a `src` folder, e.g.:

```
git clone git@github.com:TykTechnologies/tyk.git src
```

Then run the assorted task targets. For tyk rules, the location to work
on the rules is `rules/tyk/`. The taskfile bundles the official semgrep
rules and the dgryski ruleset which you can test against with `task
scan`. The reports are left in the `reports` folder.
