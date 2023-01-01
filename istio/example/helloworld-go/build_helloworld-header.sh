#!/bin/bash

set -o errexit

# Docker build variables
ENABLE_MULTIARCH_IMAGES=${ENABLE_MULTIARCH_IMAGES:-"true"}
REPO=${REPO:-"docker.io"}
PREFIX=${PREFIX:-"tanjunchen"}
VERSION=${VERSION:-"header"}

docker build -t "${REPO}/${PREFIX}/examples-helloworld-v1:${VERSION}" --build-arg service_version=header .
docker push "${REPO}/${PREFIX}/examples-helloworld-v1:${VERSION}"