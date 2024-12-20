package main

import (
	"GopherMall/user_srv/handler"
	proto "GopherMall/user_srv/proto/.UserProto"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "server ip")
	Port := flag.Int("port", 50051, "server port")

	flag.Parse()

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	err = server.Serve(lis)
	if err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}
