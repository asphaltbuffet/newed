#!/bin/bash

set -euo pipefail

rm -rf completions/*
mkdir completions

for shell in bash zsh fish; do
	echo "Generating completions for: $shell..."
	go run main.go completion "$shell" >"completions/newed.$shell"
done
