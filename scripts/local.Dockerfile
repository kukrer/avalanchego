# syntax=docker/dockerfile:experimental

# This Dockerfile is meant to be used with the build_local_dep_image.sh script
# in order to build an image using the local version of coreth

# Changes to the minimum golang version must also be replicated in
# scripts/build_savannahnode.sh
# scripts/local.Dockerfile (here)
# Dockerfile
# README.md
# go.mod
FROM golang:1.18.5-buster

RUN mkdir -p /go/src/github.com/kukrer

WORKDIR $GOPATH/src/github.com/kukrer
COPY savannahnode savannahnode
COPY coreth coreth

WORKDIR $GOPATH/src/github.com/kukrer/savannahnode
RUN ./scripts/build_savannahnode.sh
RUN ./scripts/build_coreth.sh ../coreth $PWD/build/plugins/evm

RUN ln -sv $GOPATH/src/github.com/kukrer/savannahnode-byzantine/ /savannahnonde
