#!/usr/bin/env bash
#
# E2E test for bocm provider.

# set -x # debug
set -eo errexit

TEST_DIR=`dirname "$(realpath $0)"`
ROOT_DIR="$TEST_DIR/.."

make -f "$ROOT_DIR/Makefile" build
mkdir -p "$ROOT_DIR/test/output"

# generate wechat bills output in beancount format
"$ROOT_DIR/bin/double-entry-generator" translate \
    --provider bocm \
    --config "$ROOT_DIR/example/bocm/config.yaml" \
    --output "$ROOT_DIR/test/output/test-bocm-output.beancount" \
    "$ROOT_DIR/example/bocm/example-bocm-records.csv"

diff -u --color \
    "$ROOT_DIR/example/bocm/example-bocm-output.beancount" \
    "$ROOT_DIR/test/output/test-bocm-output.beancount"

if [ $? -ne 0 ]; then
    echo "[FAIL] BOCM provider output is different from expected output."
    exit 1
fi

echo "[PASS] All BOCM provider tests!"
