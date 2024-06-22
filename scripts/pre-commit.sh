#!/bin/bash

set -euo pipefail
# set -x

mapfile -t FILES < <(git diff --cached --name-only --diff-filter=ACMR)
mapfile -t SH_FILES < <({ git diff --cached --name-only --diff-filter=ACMR && shfmt -f ./; } | sort | uniq -d)

shfmt -w -s "${SH_FILES[@]}"

golangci-lint run --new --fix

git add "${FILES[@]}"
