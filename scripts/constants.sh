#!/usr/bin/env bash
#
# Use lower_case variables in the scripts and UPPER_CASE variables for override
# Use the constants.sh for env overrides
# Use the versions.sh to specify versions
#

SAVANNAHNODE_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd ) # Directory above this script

# Set the PATHS
GOPATH="$(go env GOPATH)"
coreth_path="$GOPATH/pkg/mod/github.com/kukrer/coreth@$coreth_version"
#coreth_path="$GOPATH/src/github.com/kukrer/coreth"

# Where savannahnode binary goes
build_dir="$SAVANNAHNODE_PATH/build"
savannahnode_path="$build_dir/savannahnode"
plugin_dir="$build_dir/plugins"
evm_path="$plugin_dir/evm"

# Avalabs docker hub
# savannahlabs/savannahnode - defaults to local as to avoid unintentional pushes
# You should probably set it - export DOCKER_REPO='savannahlabs/savannahnode'
savannahnode_dockerhub_repo=${DOCKER_REPO:-"savannahnode"}

# Current branch
# TODO: fix "fatal: No names found, cannot describe anything" in github CI
current_branch=$(git symbolic-ref -q --short HEAD || git describe --tags --exact-match || true)

git_commit=${SAVANNAHNODE_COMMIT:-$( git rev-list -1 HEAD )}

# Static compilation
static_ld_flags=''
if [ "${STATIC_COMPILATION:-}" = 1 ]
then
    export CC=musl-gcc
    which $CC > /dev/null || ( echo $CC must be available for static compilation && exit 1 )
    static_ld_flags=' -extldflags "-static" -linkmode external '
fi
