FROM ubuntu:22.04

RUN apt-get -y update \
    && apt-get -y install build-essential libssl-dev git zlib1g-dev

RUN git clone https://github.com/wg/wrk.git wrk \
    && cd wrk \
    && make

RUN git clone https://github.com/giltene/wrk2.git wrk2 \
    && cd wrk2 \
    && make

RUN cp wrk/wrk /usr/local/bin
RUN mv wrk2/wrk /usr/local/bin/wrk2

HEALTHCHECK CMD exit 0

# FROM alpine:3.16 as build

# CMD ["/bin/sh"]

# WORKDIR /app

# RUN apk add --update alpine-sdk openssl-dev zlib-dev git make gcc musl-dev

# RUN apk add --no-cache git && git clone https://github.com/giltene/wrk2.git && cd wrk2 && make && mv wrk /usr/local/bin/wrk2

# RUN git clone https://github.com/wg/wrk.git && cd wrk && make && mv wrk /usr/local/bin/wrk

# ENTRYPOINT ["/bin/sh" "-c" "while sleep 864000; do :; done"]
