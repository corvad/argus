#!/bin/bash
set -euo pipefail

echo -e "running proto generation before go mod tidy so I have valid go module...\n"
./scripts/sync_proto.sh

echo -e "\nrunning go mod tidy...\n"
# Run go mod tidy with Bazel's Go toolchain
bazel run @rules_go//go -- mod tidy

echo -e "\ngo mod tidy complete."
