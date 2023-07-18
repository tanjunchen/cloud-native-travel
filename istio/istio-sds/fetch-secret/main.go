package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	discovery_v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	secret_v3 "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
)

const (
	typeSDS = "type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.Secret"

	socketPath = "/tmp/agent.sock"
)

var (
	defaultTimeout = 5 * time.Second
)

func dial(socket string) (*grpc.ClientConn, error) {
	return grpc.Dial(socket,
		grpc.WithInsecure(),
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			return (&net.Dialer{}).DialContext(ctx, "unix", addr)
		}),
		grpc.FailOnNonTempDialError(true),
		grpc.WithReturnConnectionError(),
	)
}

func main() {
	c, err := dial(socketPath)
	if err != nil {
		log.Fatal(err)
	}
	secretClient := secret_v3.NewSecretDiscoveryServiceClient(c)

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	resp, err := secretClient.FetchSecrets(ctx, &discovery_v3.DiscoveryRequest{
		TypeUrl:       typeSDS,
		ResourceNames: []string{"default"},
	})
	if err != nil {
		log.Fatal(err)
	}
	var actualSecrets []*tls_v3.Secret
	for _, resource := range resp.Resources {
		secret := new(tls_v3.Secret)
		if err := resource.UnmarshalTo(secret); err != nil {
			log.Fatal(err)
		}
		actualSecrets = append(actualSecrets, secret)
	}

	for _, v := range actualSecrets {
		fmt.Println(string(v.GetTlsCertificate().CertificateChain.GetInlineBytes()))
	}
}
