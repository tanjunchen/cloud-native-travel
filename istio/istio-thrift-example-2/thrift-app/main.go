package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sample/gen-go/Sample"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

type Greeter struct {
}

func (this *Greeter) SayHello(ctx context.Context, u *Sample.User) (r *Sample.Response, err error) {
	strJson, _ := json.Marshal(u)
	return &Sample.Response{ErrCode: 0, ErrMsg: "success", Data: map[string]string{"User": string(strJson)}}, nil
}

func (this *Greeter) GetUser(ctx context.Context, uid int32) (r *Sample.Response, err error) {
	return &Sample.Response{ErrCode: 1, ErrMsg: "user not exist."}, nil
}

func getClient(addr, protocol string, framed bool) *Sample.GreeterClient {
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
	}

	//protocol
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		fmt.Fprint(os.Stderr, "Invalid protocol specified", protocol, "\n")
		Usage()
		os.Exit(1)
	}

	//no buffered
	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()

	//framed
	if framed {
		transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	}

	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		fmt.Println("error running client:", err)
	}

	if err := transport.Open(); err != nil {
		fmt.Println("error running client:", err)
	}

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	client := Sample.NewGreeterClient(thrift.NewTStandardClient(iprot, oprot))
	return client
}

func runClient(addr, protocol string, framed bool) error {
	client := getClient(addr, protocol, framed)

	user := &Sample.User{}
	user.Name = "Hello, Thrift!"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, user)
	if err != nil {
		return err
	}

	fmt.Printf("%#v", resp)
	return nil
}

func main() {
	flag.Usage = Usage
	protocol := flag.String("P", "binary", "Specify the protocol (binary, compact, json, simplejson)")
	framed := flag.Bool("framed", false, "Use framed transport")
	buffered := flag.Bool("buffered", false, "Use buffered transport")
	addr := flag.String("addr", "localhost:9090", "Address to listen to")
	serverMode := flag.Bool("server", false, "Set to true to run as server, default to false")

	flag.Parse()

	if *serverMode == false {
		if err := runClient(*addr, *protocol, *framed); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	//protocol
	var protocolFactory thrift.TProtocolFactory
	switch *protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		fmt.Fprint(os.Stderr, "Invalid protocol specified", protocol, "\n")
		Usage()
		os.Exit(1)
	}

	//buffered
	var transportFactory thrift.TTransportFactory
	if *buffered {
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}

	//framed
	if *framed {
		transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	}

	//handler
	handler := &Greeter{}

	//transport,no secure
	var err error
	var transport thrift.TServerTransport
	transport, err = thrift.NewTServerSocket(*addr)
	if err != nil {
		fmt.Println("error running server:", err)
	}

	//processor
	processor := Sample.NewGreeterProcessor(handler)

	fmt.Println("Starting the simple server... on ", *addr)

	//start tcp server
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	err = server.Serve()

	if err != nil {
		fmt.Println("error running server:", err)
	}
}
