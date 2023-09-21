#!/bin/bash

set -o errexit

REPO=${REPO:-"docker.io"}
PREFIX=${PREFIX:-"tanjunchen"}

docker build  --platform=linux/amd64  -t "${REPO}/${PREFIX}/version:v1" --build-arg service_version=v1 .
docker push "${REPO}/${PREFIX}/version:v1"

docker build  --platform=linux/amd64  -t "${REPO}/${PREFIX}/version:v2" --build-arg service_version=v2 .
docker push "${REPO}/${PREFIX}/version:v2"

docker build  --platform=linux/amd64  -t "${REPO}/${PREFIX}/version:v3" --build-arg service_version=v3 .
docker push "${REPO}/${PREFIX}/version:v3"
