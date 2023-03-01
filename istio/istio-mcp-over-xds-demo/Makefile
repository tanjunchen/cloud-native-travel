build:
	GOOS=linux GOARCH=amd64 go build -o mcp-ovs-xds-server
	docker build -t docker.io/tanjunchen/mcp-ovs-xds-server:v1 .
	docker push docker.io/tanjunchen/mcp-ovs-xds-server:v1
	rm -rf mcp-ovs-xds-server