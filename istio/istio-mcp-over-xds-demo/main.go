package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
)

const defaultServerPort = "15010"

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer cancel()

	if err := run(ctx, os.Args); err != nil {
		log.Fatal("Error running command: ", err)
	}
}

func run(ctx context.Context, args []string) error {
	cmdName := filepath.Base(args[0])
	flags := flag.NewFlagSet(cmdName, flag.ExitOnError)
	opts, err := readOpts(flags, args)
	if err != nil {
		return fmt.Errorf("reading options: %w", err)
	}

	s := grpc.NewServer()

	adss := &adsServer{}

	discovery.RegisterAggregatedDiscoveryServiceServer(s, adss)

	l, err := net.Listen("tcp", fmt.Sprint(":", opts.serverPort))
	if err != nil {
		return fmt.Errorf("creating TCP listener: %w", err)
	}

	var g errgroup.Group
	ctx, cancel := context.WithCancel(ctx)

	log.Print("Running MCP-over-XDSv3 gRPC server")
	g.Go(func() error {
		defer cancel()
		return s.Serve(l)
	})

	g.Go(func() error {
		defer log.Print("MCP-over-XDSv3 gRPC server was shut down")
		<-ctx.Done()
		s.GracefulStop()
		return nil
	})

	t := time.NewTicker(time.Second * 30)
	defer t.Stop()

sendloop:
	for {
		select {
		case <-ctx.Done():
			adss.closeSubscribers()
			break sendloop

		case <-t.C:
			if err := adss.pushToSubscribers(); err != nil {
				log.Print("Error pushing to subscribers: ", err)
			}
		}
	}

	return g.Wait()
}

// cmdOpts are the options that can be passed to the command.
type cmdOpts struct {
	serverPort uint16
}

// readOpts parses and validates options from commmand-line flags.
func readOpts(f *flag.FlagSet, args []string) (*cmdOpts, error) {
	opts := &cmdOpts{}

	port := f.String("server-port", defaultServerPort, "Port the MCP-over-XDSv3 gRPC server listens on")

	if err := f.Parse(args[1:]); err != nil {
		return nil, err
	}

	serverPort, err := strconv.ParseUint(*port, 10, 16)
	if err != nil {
		return nil, fmt.Errorf("invalid port number: %w", err)
	}
	opts.serverPort = uint16(serverPort)

	return opts, nil
}
