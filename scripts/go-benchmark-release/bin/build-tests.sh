#!/bin/bash
releases="v5.0.15 v5.1.2 v5.2.6 v5.3.9"

# This builds test binaries for releases since v5.0-5.3, on latest patch releases.
# The binaries are *NOT* statically compiled so runtime environment is an issue.

mkdir -p tests
testFolder=$PWD/tests

# Build test binaries
for release in $releases; do
	mkdir -p cache/$release-build cache/pkg

	out=tyk-$release.test
	src=$PWD/src/tyk-$release
	srcInside=/go/pkg/github.com/TykTechnologies/tyk

	# Check out tyk if required
	if [ ! -d "$src" ]; then
		git clone --depth 1 --branch $release git@github.com:TykTechnologies/tyk.git $src
	fi

	# Build tests for release
	if [ ! -f "$testFolder/$out" ]; then
		set -x

		# Use plugin compiler as build env
		docker run --rm -v $src:$srcInside -w $srcInside -v $testFolder:/tests \
			-e GO111MODULE=on \
			-v $PWD/cache/pkg:/go/pkg \
			-v $PWD/cache/$release-build:/root/.cache/go-build \
			--entrypoint=/bin/bash tykio/tyk-plugin-compiler:$release -c "env | grep ^GO ; go mod download ; go test -o /tests/$out -trimpath -tags 'goplugin dev' -c ./gateway"

		set +x
	else
		echo $testFolder/$out OK
	fi
done

# Build master on local (no docker), otherwise we'd have to tail
# the go toolchain base image used somehow (dev dockerfile, gha?).

release=master

mkdir -p cache/$release-build cache/pkg

out=tyk-$release.test
src=$PWD/src/tyk-$release

if [ ! -d "$src" ]; then
	git clone --branch 'fix/tt-13819/benchmark-run-updates' git@github.com:TykTechnologies/tyk.git $src
fi

if [ ! -f "$testFolder/$out" ]; then
	cd $src
	CGO_ENABLED=1 go test -o $testFolder/$out -trimpath -tags 'goplugin dev' -c ./gateway
else
	echo $testFolder/$out OK
fi
