#!/bin/bash

set -o errexit

REPO=${REPO:-"docker.io"}
PREFIX=${PREFIX:-"tanjunchen"}
VERSION=${VERSION:-"latest-arm"}

docker build --no-cache=true --platform=linux/arm64 -t "${REPO}/${PREFIX}/helloworld-v1:${VERSION}" --build-arg service_version=v1 .
docker push "${REPO}/${PREFIX}/helloworld-v1:${VERSION}"