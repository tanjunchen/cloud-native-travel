FROM ubuntu:22.04 AS build-env

WORKDIR /app

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

RUN apt-get -y update \
    && apt-get -y install build-essential libssl-dev \
       git zlib1g-dev wget curl python3 libmpc-dev libmpfr-dev \
       libgmp-dev flex texinfo gperf libtool patchutils bc  libexpat-dev 

RUN git clone https://github.com/wg/wrk.git \
    && cd wrk \
    && make
RUN git clone https://github.com/giltene/wrk2.git  \
    && cd wrk2 \
    && make

FROM alpine:3.16

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

RUN apk update \
    &&  apk add bash bind-tools busybox-extras curl \
                iproute2 iputils jq mtr \
                net-tools nginx openssl \
                perl-net-telnet procps tcpdump tcptraceroute wget 

COPY --from=build-env /app/wrk/wrk /usr/local/bin/wrk
COPY --from=build-env /app/wrk2/wrk /usr/local/bin/wrk2

ENTRYPOINT ["/usr/local/bin/wrk"]
