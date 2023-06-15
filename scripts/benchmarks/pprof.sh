#!/bin/bash
if [[ -z "$2" ]]; then
	echo "Usage: $0 <name> <pprof-source>"
	exit 1
fi

name=$1
pprof=$2

if [[ ! -f "/usr/local/bin/pprofutils" ]]; then
	echo "Installing felixge/pprofutils"
	GOBIN=/usr/local/bin go install github.com/felixge/pprofutils/v2/cmd/pprofutils@latest
fi

pprofutils folded $pprof $name-folded.txt
