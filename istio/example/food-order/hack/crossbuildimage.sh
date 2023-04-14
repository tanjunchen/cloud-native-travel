#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)"
IMAGE_TAG="${IMAGE_TAG:-laest}"
#GO_LDFLAGS="$(${ROOT}/hack/version.sh)"
GO_LDFLAGS=" -X "
IMAGE_REPO_NAME="${IMAGE_REPO_NAME:-docker.io/tanjunchen}"
#PLATFORMS="${PLATFORMS:-linux/amd64,linux/arm64,linux/arm/v7}"
PLATFORMS="${PLATFORMS:-linux/amd64,linux/arm64}"

ALL_IMAGES_AND_TARGETS=(
  #{target}:{IMAGE_NAME}:{DOCKERFILE_PATH}
  food:food:dockerfile/food/Dockerfile
  order:order:dockerfile/order/Dockerfile
)

function get_imagename_by_target() {
  local key=$1
  for bt in "${ALL_IMAGES_AND_TARGETS[@]}" ; do
    local binary="${bt%%:*}"
    if [ "${binary}" == "${key}" ]; then
      local name_path="${bt#*:}"
      echo "${name_path%%:*}"
      return
    fi
  done
  echo "can not find image name: $key"
  exit 1
}

function get_dockerfile_by_target() {
  local key=$1
  for bt in "${ALL_IMAGES_AND_TARGETS[@]}" ; do
    local binary="${bt%%:*}"
    if [ "${binary}" == "${key}" ]; then
      local name_path="${bt#*:}"
      echo "${name_path#*:}"
      return
    fi
  done
  echo "can not find dockerfile for: $key"
  exit 1
}

function build_multi_arch_images() {
  local -a targets=()

  for arg in "$@"; do
    targets+=("${arg}")
  done

  if [[ ${#targets[@]} -eq 0 ]]; then
     for bt in "${ALL_IMAGES_AND_TARGETS[@]}" ; do
       targets+=("${bt%%:*}")
     done
  fi

  for arg in "${targets[@]}"; do
    IMAGE_NAME="$(get_imagename_by_target ${arg})"
    DOCKERFILE_PATH="$(get_dockerfile_by_target ${arg})"
    set -x
    # If there's any issues when using buildx, can refer to the issue below
    # https://github.com/docker/buildx/issues/495
    # https://github.com/multiarch/qemu-user-static/issues/100
    # docker run --rm --privileged multiarch/qemu-user-static --reset -p yes
    #docker buildx build --build-arg GO_LDFLAGS="${GO_LDFLAGS}" -t ${IMAGE_REPO_NAME}/${IMAGE_NAME}:${IMAGE_TAG} -f ${DOCKERFILE_PATH} --platform ${PLATFORMS} --push .
    docker buildx build --build-arg GO_LDFLAGS="" -t ${IMAGE_REPO_NAME}/${IMAGE_NAME}:${IMAGE_TAG} -f ${DOCKERFILE_PATH} --platform ${PLATFORMS} --push .
    set +x
  done
}

#use Docker Buildx to build multi-arch docker images
#How to enable Docker Buildx:
#please follow this to open Buildx function: https://medium.com/@artur.klauser/building-multi-architecture-docker-images-with-buildx-27d80f7e2408
# buildx will push the image to registry, so we need to login registry first and use `-t` flag to set image tag specifically.
build_multi_arch_images "$@"
