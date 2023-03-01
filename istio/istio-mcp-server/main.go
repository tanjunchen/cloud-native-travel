package main

import (
	"time"

	"istio-mcp-server/server"
)

func main() {
	var s = server.New(
		server.WithFreq(time.Second*1),
		server.WithAddress("0.0.0.0"),
		server.WithGRPCPort(15015),
	)
	s.Start(server.SetupSignalHandler())
}
