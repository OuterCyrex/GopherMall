package main

import (
	"GopherMall/inventory_srv/gateway"
	"GopherMall/inventory_srv/handler"
	"GopherMall/inventory_srv/initialize"
	proto "GopherMall/inventory_srv/proto/.InventoryProto"
	"GopherMall/inventory_srv/utils"
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
	proto.RegisterInventoryServer(server, handler.InventoryServer{})

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
