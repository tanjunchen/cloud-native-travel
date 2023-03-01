package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	envoycfgcorev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"

	mcpv1alpha1 "istio.io/api/mcp/v1alpha1"
	istionetv1alpha3 "istio.io/api/networking/v1alpha3"
)

var (
	maxUintDigits = len(strconv.FormatUint(uint64(math.MaxUint64), 10))
	subIDFmtStr   = `%0` + strconv.Itoa(maxUintDigits) + `d`
)

type (
	DiscoveryStream = discovery.AggregatedDiscoveryService_StreamAggregatedResourcesServer
	// DeltaDiscoveryStream is a server interface for Delta XDS.
	DeltaDiscoveryStream = discovery.AggregatedDiscoveryService_DeltaAggregatedResourcesServer
)

// adsServer implements Envoy's AggregatedDiscoveryService service for sending
// MCP resources to Istiod.
//
// See the Envoy API documentation for [Aggregated Discovery Service].
//
// Worthy resources:
// - https://pkg.go.dev/istio.io/istio/pkg/istio-agent#XdsProxy
// - https://dev.bitolog.com/grpc-long-lived-streaming/
//
// [Aggregated Discovery Service]: https://www.envoyproxy.io/docs/envoy/v1.24.0/configuration/overview/xds_api.html#aggregated-discovery-service
type adsServer struct {
	subscribers      sync.Map
	nextSubscriberID atomic.Uint64
}

// subscriber encapsulates a subscriber stream together with a unique ID.
type subscriber struct {
	id uint64

	stream      DiscoveryStream
	closeStream func()
}

var _ discovery.AggregatedDiscoveryServiceServer = (*adsServer)(nil)

func (adss *adsServer) StreamAggregatedResources(downstream DiscoveryStream) error {
	ctx, closeStream := context.WithCancel(downstream.Context())

	sub := &subscriber{
		id:          adss.nextSubscriberID.Add(1),
		stream:      downstream,
		closeStream: closeStream,
	}

	adss.subscribers.Store(sub.id, sub)

	<-ctx.Done()
	return nil
}

func (adss *adsServer) DeltaAggregatedResources(downstream DeltaDiscoveryStream) error {
	return status.Errorf(codes.Unimplemented, "not implemented")
}

// pushToSubscribers pushes MCP resources to active subscribers.
func (adss *adsServer) pushToSubscribers() error {
	mcpResources, err := makeMCPResources(numMCPResources)
	if err != nil {
		return fmt.Errorf("creating MCP resource: %w", err)
	}

	adss.subscribers.Range(func(key, value any) bool {
		log.Print("Sending to subscriber ", fmt.Sprintf(subIDFmtStr, key.(uint64)))

		if err = value.(*subscriber).stream.Send(&discovery.DiscoveryResponse{
			TypeUrl:     "networking.istio.io/v1alpha3/ServiceEntry",
			VersionInfo: strconv.FormatInt(time.Now().Unix(), 10),
			Resources:   mcpResources,
			ControlPlane: &envoycfgcorev3.ControlPlane{
				Identifier: os.Getenv("POD_NAME"),
			},
		}); err != nil {
			log.Print("Error sending MCP resources: ", err)
			value.(*subscriber).closeStream()
			adss.subscribers.Delete(key)
		}

		return true
	})

	return nil
}

// closeSubscribers closes all active subscriber streams.
func (adss *adsServer) closeSubscribers() {
	adss.subscribers.Range(func(key, value any) bool {
		log.Print("Closing stream of subscriber ", fmt.Sprintf(subIDFmtStr, key.(uint64)))
		value.(*subscriber).closeStream()
		adss.subscribers.Delete(key)

		return true
	})
}

const numMCPResources = 100

// makeMCPResources returns n Istio ServiceEntry objects serialized as protocol
// buffer messages.
func makeMCPResources(n int) ([]*anypb.Any, error) {
	mcpResources := make([]*anypb.Any, 0, numMCPResources)
	for i := 0; i < n; i++ {
		mcpRes, err := makeMCPServiceEntry(i)
		if err != nil {
			return nil, fmt.Errorf("creating MCP resource: %w", err)
		}
		mcpResources = append(mcpResources, mcpRes)
	}

	return mcpResources, nil
}

// makeMCPServiceEntry returns an Istio ServiceEntry serialized as a protocol
// buffer message.
func makeMCPServiceEntry(idx int) (*anypb.Any, error) {
	seSpec := &istionetv1alpha3.ServiceEntry{
		Hosts:    []string{fmt.Sprintf("fake%03d.example.com", idx)},
		Location: istionetv1alpha3.ServiceEntry_MESH_EXTERNAL,
		Ports: []*istionetv1alpha3.ServicePort{{
			Number:   443,
			Name:     "https",
			Protocol: "TLS",
		}},
		Resolution: istionetv1alpha3.ServiceEntry_STATIC,
		Endpoints: []*istionetv1alpha3.WorkloadEntry{{
			// TEST-NET-1
			// https://datatracker.ietf.org/doc/html/rfc5737
			Address: "192.0.2.42",
		}},
	}

	mcpResBody := &anypb.Any{}
	if err := anypb.MarshalFrom(mcpResBody, seSpec, proto.MarshalOptions{}); err != nil {
		return nil, fmt.Errorf("serializing ServiceEntry to protobuf message: %w", err)
	}

	mcpResTyped := &mcpv1alpha1.Resource{
		Metadata: &mcpv1alpha1.Metadata{
			Name: fmt.Sprintf("istio-system/mcp-example-%03d", idx),
		},
		Body: mcpResBody,
	}

	mcpRes := &anypb.Any{}
	if err := anypb.MarshalFrom(mcpRes, mcpResTyped, proto.MarshalOptions{}); err != nil {
		return nil, fmt.Errorf("serializing MCP Resource to protobuf message: %w", err)
	}

	return mcpRes, nil
}
