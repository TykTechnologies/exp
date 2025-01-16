# Release benchmarks

The scripts contained within run go benchmarks against the latest
5.0-5.3 patch releases. The versions are inlined in `bin/*.sh`.

Plugin compiler generally contains the build environment for each
particular release. This means that a tagged plugin compiler will have
the complete gateway source code and the go toolchain used to build the
release and in our case, tests. However, the LTS branch (or master) is
a fluid codebase - so there's some `git clone` under the hood.

For each release, we:

- Check out a gateway source tree
- Build the test binary in `tests/`
- Run docker services from latest gateway
- Run each benchmark individually
- Collect cpu, memory and trace profiles (disabled currently)
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

# Running everything

task: Available tasks for this project:

* build:        Build test binaries
* clean:        Clean up failed tests
* readme:       Update readme if needed
* report:       Print test run durations CSV
* run:          Run test binaries

## task: build

Build test binaries

commands:
 - ./bin/build-tests.sh

## task: clean

Clean up failed tests

commands:
 - sudo find out -name '*.log' -size -140c -delete

## task: readme

Update readme if needed

commands:
 - sh ./README.md.sh > README.md

## task: report

Print test run durations CSV

commands:
 - php bin/report.php

## task: run

Run test binaries

commands:
 - ./bin/run-tests.sh

