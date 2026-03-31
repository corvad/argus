#!/bin/bash
set -euo pipefail

echo -e "generating proto files...\n"

# Find all .proto files and derive their Bazel targets
for proto in $(find proto -name "*.proto"); do
    dir=$(dirname "$proto")
    pkg=$(basename "$dir")
    
    bazel build //"$dir":"${pkg}_go_proto"
    find "bazel-bin/$dir" -name "*.pb.go" -exec cp -f {} "$dir/" \;
done

echo -e "\nproto files synced."
