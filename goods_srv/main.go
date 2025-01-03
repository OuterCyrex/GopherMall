package main

import (
	"GopherMall/goods_srv/gateway"
	"GopherMall/goods_srv/handler"
	"GopherMall/goods_srv/initialize"
	proto "GopherMall/goods_srv/proto/.GoodsProto"
	"GopherMall/goods_srv/utils"
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
	proto.RegisterGoodsServer(server, handler.GoodsServer{})

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
