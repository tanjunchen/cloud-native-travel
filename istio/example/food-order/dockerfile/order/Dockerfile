FROM golang:1.18.3-alpine3.15 AS builder

ARG GO_LDFLAGS

WORKDIR /go/src/servicemesh-example

RUN CGO_ENABLED=0 GO111MODULE=off go build -v -ldflags "${GO_LDFLAGS} -w -s" -o /usr/local/bin/order /go/src/servicemesh-example/order/cmd/main.go

FROM alpine:3.13

COPY --from=builder /usr/local/bin/order /usr/local/bin/order

ENTRYPOINT ["order"]