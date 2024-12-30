package main

import (
	"GopherMall/user_srv/gateway"
	"GopherMall/user_srv/handler"
	"GopherMall/user_srv/initialize"
	proto "GopherMall/user_srv/proto/.UserProto"
	"GopherMall/user_srv/utils"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"runtime"
)

func main() {
	isDebug := runtime.GOOS == "windows"

	initialize.InitLogger()
	initialize.InitConfig(isDebug)
	initialize.InitMysql()

	// 动态端口分配
	port, err := utils.GetFreePort()
	if err != nil {
		zap.S().Panicf("Get Free Port error:%v", err)
	}

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})

	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		zap.S().Panicf("failed to listen: %v", err)
	}

	zap.S().Infof("Server Runs On Port %d", port)

	//注册健康检查供consul使用
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	//健康检查
	gateway.HealthCheck(fmt.Sprintf("127.0.0.1:%d", port), 15)

	err = server.Serve(lis)
	if err != nil {
		zap.S().Panicf("failed to serve: %v", err)
	}
}
