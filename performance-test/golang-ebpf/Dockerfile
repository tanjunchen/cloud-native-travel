FROM golang:1.20.2-alpine3.16 AS build-env
ADD . /src
ENV GOPROXY https://goproxy.cn
RUN cd /src && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o goapp
#RUN cd /src && CGO_ENABLED=0 go build -o goapp

FROM ubuntu:xenial
RUN apt-get -y update
RUN apt-get --yes install libav-tools bash  automake autoconf  make 
RUN apt-get -y install curl wget net-tools

WORKDIR /app
COPY --from=build-env /src/goapp /app/

ENTRYPOINT ["./goapp"]
