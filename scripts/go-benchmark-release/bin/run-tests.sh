#!/bin/bash
releases="v5.0.15 v5.1.2 v5.2.6 v5.3.9"

# Start required test services (redis, httpbin...)
function servicesUp {
	src=$PWD/src/tyk-master
	cd $src && task services:up && cd -
}

# Shut down test services
function servicesDown {
	src=$PWD/src/tyk-master
	rm -rf $PWD/github.com
	cd $src && task services:down && cd -
}

# Run benchmarks for version
function benchmark {
	local version=$1
	local BENCHMARKS=$(./tests/tyk-$version.test -test.list=Bench.+)
	for name in $BENCHMARKS; do
		if [ -f "$version/$name.log" ]; then
			echo "Skipping $name, exists"
			continue
		fi

		# This is a very slow benchmark. We skip it since it timeouts.
		if [[ "$name" == "BenchmarkPurgeLapsedOAuthTokens" ]]; then
			continue
		fi

		# Work around runtime.Caller and trimpath
		rm -rf $PWD/github.com
		mkdir -p $PWD/github.com/TykTechnologies/
		ln -sf $PWD/src/tyk-$version github.com/TykTechnologies/tyk

		local output=out/$version
		mkdir -p $output

		if [ -f "$output/$name.log" ]; then
			echo "$output/$name OK"
			continue
		fi

		set -x
		timeout -k 120s 60s \
			./tests/tyk-$version.test -test.run=^$ \
				-test.bench=^${name}$ \
				-test.benchtime 15s \
				2>$output/$name.err | tee $output/$name.log
		set +x

		# This test invocation collects more details.
		# It also fails more.

#		timeout -k 120s 60s \
#			./tests/tyk-$version.test -test.run=^$ \
#				-test.bench=^${name}$ \
#				-test.benchtime 30s -test.benchmem \
#				-test.cpuprofile=$output/$name-cpu.out \
#				-test.memprofile=$output/$name-mem.out \
#				-test.trace=$output/$name-trace.out \
#				2>$output/$name.err | tee $output/$name.log

		exitCode=$?
		if (( $exitCode > 0 )); then
			echo "Benchmark failure, exit: $exitCode" > /dev/stderr
		fi

		if [ -f "$output/$name.err" ]; then
			cat $output/$name.err > /dev/stderr
		fi

		sleep 1
	done
}

trap "servicesDown" EXIT

servicesUp

if [ ! -z "$1" ]; then
	benchmark $1
else
	benchmark master
	for release in $releases; do
		benchmark $release
	done
fi
