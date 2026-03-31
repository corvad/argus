#!/bin/bash

# Copyright (c) 2026 David Corvaglia. All rights reserved.
# Use of this source code is governed by an MIT license
# that can be found in the LICENSE file.

set -euo pipefail

echo -e "running proto generation before go mod tidy so I have valid go module...\n"
./scripts/sync_proto.sh

echo -e "\nrunning go mod tidy...\n"
# Run go mod tidy with Bazel's Go toolchain
bazel run @rules_go//go -- mod tidy

echo -e "\ngo mod tidy complete."
