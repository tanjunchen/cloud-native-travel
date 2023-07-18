package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"

	pb "grpcdemo/protos"
)

func main() {
	// 监听本地端口
	listener, err := net.Listen("tcp", ":10001")
	if err != nil {
		return
	}
	s := grpc.NewServer()                      // 创建GRPC
	pb.RegisterEchoServiceServer(s, &server{}) // 在GRPC服务端注册服务

	reflection.Register(s)
	fmt.Println("grpc server 10001")
	if err = s.Serve(listener); err != nil {
		log.Fatalln("failed to serve ", err)
	}
}

type server struct{}

func (s *server) Echo(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	ip := GetRealAddr(ctx)
	fmt.Println(ip, request.Message)
	res := "echo() -> ip [ " + ip + " ] param [ " + request.Message + " ]"
	return &pb.HelloResponse{Message: res}, nil
}

// GetRealAddr get real client ip
func GetRealAddr(ctx context.Context) string {
	var addr string
	if pr, ok := peer.FromContext(ctx); ok {
		if tcpAddr, ok := pr.Addr.(*net.TCPAddr); ok {
			addr = tcpAddr.IP.String()
		} else {
			addr = pr.Addr.String()
		}
	}
	return addr
}
