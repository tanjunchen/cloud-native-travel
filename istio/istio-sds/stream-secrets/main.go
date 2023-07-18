package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"

	"google.golang.org/grpc"

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

func waitForCtrlC() {
	var wg sync.WaitGroup
	wg.Add(1)
	var signalCh chan os.Signal
	signalCh = make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	go func() {
		<-signalCh
		wg.Done()
	}()
	wg.Wait()
}
func main() {
	c, err := dial(socketPath)
	if err != nil {
		log.Fatal(err)
	}
	secretClient := secret_v3.NewSecretDiscoveryServiceClient(c)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	req := &discovery_v3.DiscoveryRequest{
		TypeUrl:       typeSDS,
		ResourceNames: []string{"default"},
	}
	stream, err := secretClient.StreamSecrets(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.CloseSend()
	stream.Send(req)

	updateChan := make(chan discovery_v3.DiscoveryResponse, 1)

	go func() {
		for {
			msq, err := stream.Recv()
			if err != nil {
				log.Fatal(err)
			}
			updateChan <- *msq
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case resp := <-updateChan:
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
		case <-sigs:
			close(updateChan)
			log.Fatal("stop updates")
		}
	}
}
