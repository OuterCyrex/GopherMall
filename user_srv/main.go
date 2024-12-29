package main

import (
	"GopherMall/user_srv/gateway"
	"GopherMall/user_srv/handler"
	"GopherMall/user_srv/initialize"
	proto "GopherMall/user_srv/proto/.UserProto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

func main() {

	initialize.InitLogger()
	initialize.InitConfig(true)
	initialize.InitMysql()

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		zap.S().Panicf("failed to listen: %v", err)
	}

	//注册健康检查供consul使用
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	//健康检查
	gateway.HealthCheck()

	err = server.Serve(lis)
	if err != nil {
		zap.S().Panicf("failed to serve: %v", err)
	}
}
