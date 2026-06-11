# Promote packages

This workflow publishes an existing package version from Packagecloud unstable
repositories into the corresponding stable repositories.

This workflow replaces the previous Buddy `Promote packages` pipeline.

## How to run it

Run the `Promote packages` workflow manually from GitHub Actions.

Required inputs:

- `repo`: Packagecloud repository to promote from, for example `tyk-gateway`.
- `version`: Package version to promote, for example `5.8.0`.

Optional inputs:

- `debvers`: Space-separated Debian/Ubuntu distributions.
- `rpmvers`: Space-separated RPM distributions.
- `dry_run`: Print the Packagecloud commands without changing Packagecloud.

At least one of `debvers` or `rpmvers` must be set.

## Dry run

Use `dry_run` to review the exact Packagecloud actions before changing
Packagecloud.

When `dry_run` is `true`, the workflow prints the package names, source repos,
destination repos, and commands it would run. It does not yank, promote, push,
download, or install the Packagecloud CLI.

This is useful before the first real run, or when checking a new version or
distribution list.

## What it promotes

For Debian/Ubuntu distributions, the workflow promotes both package arches:

- `amd64`
- `arm64`

For RPM distributions, the workflow promotes both package arches:

- `x86_64`
- `aarch64`

For a normal repository such as `tyk-gateway`, the workflow promotes packages
from:

```text
tyk/<repo>-unstable/<distro>
```

to:

```text
tyk/<repo>/<distro>
```

For example:

```text
tyk/tyk-gateway-unstable/ubuntu/jammy
tyk/tyk-gateway/ubuntu/jammy
```

## Package names

Debian package names are generated as:

```text
<repo>_<version>_<arch>.deb
```

Example:

```text
tyk-gateway_5.8.0_amd64.deb
```

RPM package names are generated as:

```text
<repo>-<version>-1.<arch>.rpm
```

Example:

```text
tyk-gateway-5.8.0-1.x86_64.rpm
```
