#!/bin/bash

set -o errexit

# Docker build variables
ENABLE_MULTIARCH_IMAGES=${ENABLE_MULTIARCH_IMAGES:-"true"}
REPO=${REPO:-"docker.io"}
PREFIX=${PREFIX:-"tanjunchen"}
VERSION=${VERSION:-"latest"}

if [ "${ENABLE_MULTIARCH_IMAGES}" == "true" ]; then
  PLATFORMS="linux/arm64,linux/amd64"
  DOCKER_BUILD_ARGS="docker buildx build --platform ${PLATFORMS} --push"
  # Install QEMU emulators
  docker run --rm --privileged tonistiigi/binfmt --install all
  docker buildx rm multi-builder || :
  docker buildx create --use --name multi-builder --platform ${PLATFORMS}
  docker buildx use multi-builder
else
  DOCKER_BUILD_ARGS="docker build"
fi

${DOCKER_BUILD_ARGS} --pull -t "${REPO}/${PREFIX}/examples-helloworld-v1:${VERSION}" --build-arg service_version=v1 .
${DOCKER_BUILD_ARGS} --pull -t "${REPO}/${PREFIX}/examples-helloworld-v2:${VERSION}" --build-arg service_version=v2 .
