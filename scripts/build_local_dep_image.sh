#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

echo "Building docker image based off of most recent local commits of savannahnode and coreth"

SAVANNAHNODE_REMOTE="git@github.com:kukrer/savannahnode.git"
CORETH_REMOTE="git@github.com:kukrer/coreth.git"
DOCKERHUB_REPO="savannahlabs/savannahnode"

DOCKER="${DOCKER:-docker}"
SCRIPT_DIRPATH=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)
ROOT_DIRPATH="$(dirname "${SCRIPT_DIRPATH}")"

AVA_LABS_RELATIVE_PATH="src/github.com/kukrer"
EXISTING_GOPATH="$GOPATH"

export GOPATH="$SCRIPT_DIRPATH/.build_image_gopath"
WORKPREFIX="$GOPATH/src/github.com/kukrer"

# Clone the remotes and checkout the desired branch/commits
SAVANNAHNODE_CLONE="$WORKPREFIX/savannahnode"
CORETH_CLONE="$WORKPREFIX/coreth"

# Replace the WORKPREFIX directory
rm -rf "$WORKPREFIX"
mkdir -p "$WORKPREFIX"


SAVANNAHNODE_COMMIT_HASH="$(git -C "$EXISTING_GOPATH/$AVA_LABS_RELATIVE_PATH/savannahnode" rev-parse --short HEAD)"
CORETH_COMMIT_HASH="$(git -C "$EXISTING_GOPATH/$AVA_LABS_RELATIVE_PATH/coreth" rev-parse --short HEAD)"

git config --global credential.helper cache

git clone "$SAVANNAHNODE_REMOTE" "$SAVANNAHNODE_CLONE"
git -C "$SAVANNAHNODE_CLONE" checkout "$SAVANNAHNODE_COMMIT_HASH"

git clone "$CORETH_REMOTE" "$CORETH_CLONE"
git -C "$CORETH_CLONE" checkout "$CORETH_COMMIT_HASH"

CONCATENATED_HASHES="$SAVANNAHNODE_COMMIT_HASH-$CORETH_COMMIT_HASH"

"$DOCKER" build -t "$DOCKERHUB_REPO:$CONCATENATED_HASHES" "$WORKPREFIX" -f "$SCRIPT_DIRPATH/local.Dockerfile"
