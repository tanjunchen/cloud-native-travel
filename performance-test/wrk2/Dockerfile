FROM alpine:3.16 as build
#ARG WRK2_COMMIT_HASH=44a94c17d8e6a0bac8559b53da76848e430cb7a7
# Install deps and build from source
RUN apk add --no-cache openssl-dev zlib-dev git make gcc musl-dev
RUN git clone https://github.com/giltene/wrk2 && cd wrk2 && make


FROM alpine:3.16
# it seems libgcc_s.so is dynamically linked so we need to install libgcc
RUN apk add --no-cache libgcc
#RUN adduser wrk_user -D -H
#USER wrk_user

COPY --from=build /wrk2/wrk /usr/bin/wrk2
ENTRYPOINT ["/usr/bin/wrk2"]

