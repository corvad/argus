#!/bin/bash

# Copyright (c) 2026 David Corvaglia. All rights reserved.
# Use of this source code is governed by an MIT license
# that can be found in the LICENSE file.

set -euo pipefail

YEAR=$(date +%Y)
OWNER="${1:?Usage: $0 <owner>}"

SLASH_HEADER="// Copyright (c) $YEAR $OWNER. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

"

HASH_HEADER="# Copyright (c) $YEAR $OWNER. All rights reserved.
# Use of this source code is governed by an MIT license
# that can be found in the LICENSE file.

"

add_header() {
    local file="$1"
    local header="$2"

    if head -5 "$file" | grep -q "Copyright"; then
        echo "Skipping $file (already has header)"
        return
    fi

    local tmp
    tmp=$(mktemp)

    if head -1 "$file" | grep -q "^#!"; then
        shebang=$(head -1 "$file")
        rest=$(tail -n +2 "$file")
        printf '%s\n\n%s%s\n' "$shebang" "$header" "$rest" > "$tmp"
    else
        printf '%s%s\n' "$header" "$(cat "$file")" > "$tmp"
    fi

    mv "$tmp" "$file"
    echo "Added header to $file"
}

# Go files (exclude generated)
while read -r f; do
    add_header "$f" "$SLASH_HEADER"
done < <(find . -name "*.go" -not -path "./bazel-*" -not -name "*.pb.go")

# Proto files
while read -r f; do
    add_header "$f" "$SLASH_HEADER"
done < <(find . -name "*.proto" -not -path "./bazel-*")

# Shell scripts (exclude this script)
SELF="$(basename "$0")"
while read -r f; do
    add_header "$f" "$HASH_HEADER"
done < <(find . -name "*.sh" -not -path "./bazel-*" -not -name "$SELF")

echo "Done."