#!/bin/bash
set -e

if [ -z "$GITHUB_TAG" ]; then
    echo "Need GITHUB_TAG env"
    exit 1
fi

export PLUGIN_SOURCE_PATH=./basic-plugin
export TYK_GW_PATH=${TYK_GW_PATH:-$(git rev-parse --show-toplevel)}

./build.sh plugin.so