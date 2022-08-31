#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Savannahnode root folder
SAVANNAHNODE_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )
# Load the versions
source "$SAVANNAHNODE_PATH"/scripts/versions.sh
# Load the constants
source "$SAVANNAHNODE_PATH"/scripts/constants.sh

# Download dependencies
echo "Downloading dependencies..."
go mod download

# Build savannahnode
"$SAVANNAHNODE_PATH"/scripts/build_savannahnode.sh

# Build coreth
"$SAVANNAHNODE_PATH"/scripts/build_coreth.sh

# Exit build successfully if the binaries are created
if [[ -f "$savannahnode_path" && -f "$evm_path" ]]; then
        echo "Build Successful"
        exit 0
else
        echo "Build failure" >&2
        exit 1
fi
