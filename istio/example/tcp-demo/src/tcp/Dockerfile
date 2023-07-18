ARG BASE_DISTRIBUTION=default

# build a tcp-echo binary using the golang container
FROM golang:1.13 as builder

WORKDIR /tanjunchen/istio/tcp-echo-server/

COPY server/main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tcp-example main.go

FROM scratch as default

FROM ${BASE_DISTRIBUTION}

WORKDIR /bin/

COPY --from=builder /tanjunchen/istio/tcp-echo-server/tcp-example .

ENTRYPOINT [ "/bin/tcp-example" ]

CMD [ "9000", "tcp" ]

EXPOSE 9000
