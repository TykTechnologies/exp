# workflow-lint

`workflow-lint` is a tool that checks if all `uses` lines in GitHub Actions workflow files reference the same version of an action. It helps ensure consistency across workflows and can automatically update actions to the latest versions, including the ability to skip specific files during analysis.

## Features

- Detects inconsistent versions of GitHub Actions across workflow files.
- Automatically updates actions to the latest major versions.
- Allows for user-defined action version overrides.
- Supports skipping specific workflow files from analysis.

## Installation

You can compile the Go program by running:

```bash
go build -o workflow-lint main.go
```

## Usage

Run `workflow-lint` in your repository to check for version mismatches and update GitHub Actions in your workflow files.

### Basic Command

```bash
./workflow-lint
```

By default, this command will analyze all workflow files in the `.github/workflows/` directory and check for version mismatches.

### Flags

- `--fix`: Automatically update actions to their latest versions.
  
  Example:
  ```bash
  ./workflow-lint --fix
  ```

- `--list`: List the latest action versions used in the workflows.
  
  Example:
  ```bash
  ./workflow-lint --list
  ```

- `--base`: Specify a comma-separated list of `action@version` pairs to override the default versions.
  
  Example:
  ```bash
  ./workflow-lint --base=actions/checkout@v4,actions/setup-node@v4
  ```

- `--ignore`: Specify a comma-separated list of workflow files to skip during analysis.
  
  Example:
  ```bash
  ./workflow-lint --fix --ignore=release.yml,test.yml
  ```

### Default Action Versions

By default, `workflow-lint` checks and updates actions to the following versions:

- `actions/setup-node@v4`
- `actions/cache@v4`
- `actions/setup-go@v5`
- `actions/download-artifact@v4`
- `actions/checkout@v4`
- `actions/setup-python@v5`
- `actions/upload-artifact@v4`

You can override these default versions using the `--base` flag.
