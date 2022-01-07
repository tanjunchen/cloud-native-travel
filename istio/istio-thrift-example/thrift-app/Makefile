REPO ?= tanjunchen/thrift-example
VERSION ?= v1.0
IMAGE=$(REPO):$(VERSION)

default: all

generate:
	docker run --rm -v `pwd`:/data -it thrift thrift --gen go  -r -o /data  /data/user.thrift
	docker run --rm -v `pwd`:/data -it thrift thrift --gen go  -r -o /data  /data/server.thrift

image:
	docker build -t $(IMAGE) .

push:
	docker push $(IMAGE)

all: image push