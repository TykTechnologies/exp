#!/bin/bash
if [[ -z "$2" ]]; then
	echo "Usage: $0 <name> <url>"
	exit 1
fi

name=$1
url=$2

if [[ ! -f "/usr/local/bin/hey-hdr" ]]; then
	echo "Installing asoorm/hey-hdr"
	GOBIN=/usr/local/bin go install github.com/asoorm/hey-hdr@latest
fi

if [[ ! -f "/usr/local/bin/pprofutils" ]]; then
	echo "Installing felixge/pprofutils"
	GOBIN=/usr/local/bin go install github.com/felixge/pprofutils/v2/cmd/pprofutils@latest
fi

time hey -n 5000 -c 10 -H 'authorization: f67ba7e5805f4992585508b7e58d7c6c' -o csv $url > $name-http-client.csv
cat $name-http-client.csv | hey-hdr -out $name
