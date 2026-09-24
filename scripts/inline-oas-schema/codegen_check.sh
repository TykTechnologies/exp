#!/usr/bin/env bash
# Generates Python, TypeScript, Java and Go clients from a swagger and builds them.
#
# Python and TypeScript must be clean, and a Tyk OAS API model must round-trip
# openapi/info/paths/x-tyk-api-gateway. Java and Go fail only on errors in the
# inlined OAS3* models; other errors come from the source swagger (TT-18474) and
# are reported as warnings. Set STRICT=1 to fail on any error.
#
# Usage: codegen_check.sh <swagger.yml>
set -uo pipefail

SPEC=$(realpath "$1")
HERE=$(cd "$(dirname "$0")" && pwd)
GENERATOR=openapitools/openapi-generator-cli:v7.10.0
WORK=$(mktemp -d)
FAILED=0

fail() { echo "::error::$1"; FAILED=1; }
warn() { if [ "${STRICT:-0}" = 1 ]; then fail "$1"; else echo "::warning::$1"; fi; }

generate() {
  docker run --rm --network none -u "$(id -u):$(id -g)" -v "$(dirname "$SPEC"):/spec:ro" -v "$WORK:/out" \
    "$GENERATOR" generate -i "/spec/$(basename "$SPEC")" -g "$1" -o "/out/$1" > "$WORK/$1.log" 2>&1 \
    || { fail "$1: generation failed"; tail -20 "$WORK/$1.log"; return 1; }
}

# Errors in the inlined OAS3* model files are ours; the rest come from the source swagger.
split_errors() {
  local log=$1 lang=$2 line=$3 ours_file=$4
  local ours upstream
  ours=$(grep -E "$line" "$log" | grep -cE "$ours_file")
  upstream=$(grep -E "$line" "$log" | grep -cvE "$ours_file")
  [ "$ours" -eq 0 ] && [ "$upstream" -eq 0 ] && { fail "$lang: build failed without compiler errors"; tail -20 "$log"; }
  [ "$ours" -gt 0 ] && { fail "$lang: $ours compile errors in inlined OAS3* models"; grep -E "$line" "$log" | grep -E "$ours_file" | head -20; }
  [ "$upstream" -gt 0 ] && warn "$lang: $upstream compile errors from the source swagger (TT-18474)"
  return 0
}

echo "== Python"
if generate python; then
  python3 -m venv "$WORK/venv" && "$WORK/venv/bin/pip" install --quiet pydantic python-dateutil urllib3 typing-extensions
  (cd "$WORK/python" && PYTHONPATH=. "$WORK/venv/bin/python" "$HERE/codegen_roundtrip.py") || fail "python: import or round-trip failed"
fi

echo "== TypeScript"
if generate typescript-axios; then
  (cd "$WORK/typescript-axios" && npm install --silent --no-save axios typescript@5 > /dev/null 2>&1 \
    && npx tsc --noEmit --skipLibCheck --target es2019 --moduleResolution node ./*.ts > "$WORK/ts.log" 2>&1) \
    || { fail "typescript: $(grep -c 'error TS' "$WORK/ts.log") compile errors"; head -20 "$WORK/ts.log"; }
fi

echo "== Java"
if generate java; then
  (cd "$WORK/java" && mvn -q -B -DskipTests -Dmaven.compiler.maxerrs=1000 compile > "$WORK/java.log" 2>&1) \
    || split_errors "$WORK/java.log" java '^\[ERROR\] /' '/model/OAS3[A-Za-z]*\.java'
fi

echo "== Go"
if generate go; then
  (cd "$WORK/go" && go mod tidy > /dev/null 2>&1 && go build -gcflags=-e ./... > "$WORK/go.log" 2>&1) \
    || split_errors "$WORK/go.log" go '^\./' '^\./model_oas3_'
fi

rm -rf "$WORK"
[ "$FAILED" = 0 ] && echo "✅ Codegen check passed for $SPEC" || echo "❌ Codegen check failed for $SPEC"
exit "$FAILED"
