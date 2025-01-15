# Release benchmarks

The scripts contained within run go benchmarks against the latest
5.0-5.3 patch releases. The versions are inlined in `bin/*.sh`.

Plugin compiler generally contains the build environment for each
particular release. This means that a tagged plugin compiler will have
the complete gateway source code and the go toolchain used to build the
release and in our case, tests.

- Check out a gateway source tree
- Build the test binary in `tests/`
- Run docker services from latest gateway
- Run each benchmark individually
- Collect cpu, memory and trace profiles
- Collect benchmark outputs

When all the benchmarks are collected, CSV output is produced using task
report, which can be directly copy pasted into a google spreadsheet.

## Requirements

Go, git, php (php-cli) until someone ports the report script to
something else. Or runs php:alpine or something via docker.

## Guide

Just paste the output produced by `task report` into a google sheets
doc. For implementation info see the following JIRA tickets:

- https://tyktech.atlassian.net/browse/TT-13017
- https://tyktech.atlassian.net/browse/TT-13819
