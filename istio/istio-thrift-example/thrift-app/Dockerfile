FROM golang

WORKDIR /thrift-app

COPY . /thrift-app

ARG TARGET=example

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64                      \
	go build -i -v -o /tmp/${TARGET} .

FROM alpine

ARG TARGET=example

COPY --from=0 /tmp/${TARGET} /${TARGET}

RUN ln -s /${TARGET} /entrypoint

ENTRYPOINT ["/entrypoint"]
