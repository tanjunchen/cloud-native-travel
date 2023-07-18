package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "grpcdemo/protos"
)

const (
	URL = "provider-demo"
	//URL = "localhost:10001"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Consumer Demo")
	})
	e.GET("/echo-rest/:msg", getMsg)
	e.Logger.Fatal(e.Start(":9999"))
}

func getMsg(c echo.Context) error {
	msg := c.Param("msg")
	if msg == "" {
		return c.String(http.StatusBadRequest, "no parameter")
	}
	md := metadata.New(getForwardHeaders(c.Request()))
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(URL, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Can't connect: " + URL)
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)
	resp, err := client.Echo(ctx, &pb.HelloRequest{
		Message: msg,
	})
	if err != nil {
		log.Fatalln("Do Format error:" + err.Error())
	}
	log.Println("=====>", resp.Message)
	return c.String(http.StatusOK, resp.Message)
}

func getForwardHeaders(r *http.Request) map[string]string {
	headers := make(map[string]string)
	forwardHeaders := []string{
		"user",
		"x-request-id",
		"x-b3-traceid",
		"x-b3-spanid",
		"x-b3-parentspanid",
		"x-b3-sampled",
		"x-b3-flags",
		"x-ot-span-context",
	}

	for _, h := range forwardHeaders {
		if v := r.Header.Get(h); v != "" {
			headers[h] = v
		}
	}

	return headers
}
