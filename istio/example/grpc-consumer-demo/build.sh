#!/bin/bash
#cd $WORKSPACE
export GOPROXY=https://goproxy.io

 #根据 go.mod 文件来处理依赖关系。
go mod tidy

# linux环境编译
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goapp
# 构建docker镜像，项目中需要在当前目录下有dockerfile，否则构建失败

docker build  -t bms-golang-grpc-consumer-demo:v1 .

docker tag  bms-golang-grpc-consumer-demo:v1   106.12.174.224/public/bms-golang-grpc-consumer-demo:v1
docker push 106.12.174.224/public/bms-golang-grpc-consumer-demo:v1
docker rmi  bms-golang-grpc-consumer-demo:v1
docker rmi  106.12.174.224/public/bms-golang-grpc-consumer-demo:v1