package main

import (
	"GopherMall/userop_srv/gateway"
	"GopherMall/userop_srv/handler"
	"GopherMall/userop_srv/initialize"
	AddressProto "GopherMall/userop_srv/proto/.AddressProto"
	FavProto "GopherMall/userop_srv/proto/.FavProto"
	MessageProto "GopherMall/userop_srv/proto/.MessageProto"
	"GopherMall/userop_srv/utils"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitMysql()

	server := grpc.NewServer()
	AddressProto.RegisterAddressServer(server, handler.AddressServer{})
	MessageProto.RegisterMessageServer(server, handler.MessageServer{})
	FavProto.RegisterFavServer(server, handler.FavServer{})

	port, err := utils.GetFreePort()
	if err != nil {
		zap.S().Panicf("Get Free Port error:%v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		zap.S().Panicf("dial tcp failed: %v", err)
	}

	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	gateway.HealthCheck(fmt.Sprintf("127.0.0.1:%d", port), 14)

	zap.S().Infof("grpc server Runs On: %d", port)

	err = server.Serve(lis)
	if err != nil {
		zap.S().Panicf("Run grpc server failed: %v", err)
	}
}
