#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Directory above this script
SAVANNAHNODE_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )

# Load the versions
source "$SAVANNAHNODE_PATH"/scripts/versions.sh

# Load the constants
source "$SAVANNAHNODE_PATH"/scripts/constants.sh

# WARNING: this will use the most recent commit even if there are un-committed changes present
full_commit_hash="$(git --git-dir="$SAVANNAHNODE_PATH/.git" rev-parse HEAD)"
commit_hash="${full_commit_hash::8}"

echo "Building Docker Image with tags: $savannahnode_dockerhub_repo:$commit_hash , $savannahnode_dockerhub_repo:$current_branch"
docker build -t "$savannahnode_dockerhub_repo:$commit_hash" \
        -t "$savannahnode_dockerhub_repo:$current_branch" "$SAVANNAHNODE_PATH" -f "$SAVANNAHNODE_PATH/Dockerfile"
