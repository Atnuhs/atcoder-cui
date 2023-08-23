#!/bin/bash

set -eu

url=$1

readonly DIR="$(dirname $0)/${url##*/}"
mkdir -p "${DIR}"
file="${DIR}/verify.test.go"
main="$(dirname $0)/../main.go"

{
    echo "// verification-helper: PROBLEM ${url}"
    cat "$main"
} | tee "$file"

echo "new file created: $file"
