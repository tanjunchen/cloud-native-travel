#!/bin/bash

set -e

# bash build.sh docker.io/tanjunchen/bms-golang-grpc-consumer-demo:v1

export GOPROXY=https://goproxy.io
go mod tidy
# 镜像地址信息
IMAGE=$1
SHELL_FOLDER=$(cd "$(dirname "$0")"; pwd)
docker build -t ${IMAGE} .
docker push ${IMAGE}
