package initialize

import (
	"GopherMall/user_api/gateway"
	"GopherMall/user_api/global"
	proto "GopherMall/user_srv/proto/.UserProto"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// InitSrvConnection 用于实现负载均衡
func InitSrvConnection(wait uint, policy string) {
	userConn, err := grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=%ds",
			global.ServerConfig.Consul.Host,
			global.ServerConfig.Consul.Port,
			global.ServerConfig.UserSrv.Name,
			wait,
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"loadBalancingPolicy": "%s"}`, policy)),
	)

	if err != nil {
		zap.S().Panicf("Load Balance Init Failed: %v", err)
		return
	}

	global.UserSrvClient = proto.NewUserClient(userConn)
}

// Deprecated: Use InitSrvConnection instead.
func InitSrvConnection2() {

	// 从Consul中拉取 User_srv 服务信息
	data, err := gateway.PullServiceByName(global.ServerConfig.UserSrv.Name)
	if err != nil {
		zap.S().Panicf("Init Service Connection Failed: %v", err)
		return
	}

	// grpc 连接对应服务
	userConn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d",
			data[global.ServerConfig.UserSrv.Name].Address,
			data[global.ServerConfig.UserSrv.Name].Port,
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		zap.S().Panicf("Connect to Grpc Server Failed: %v", err)
		return
	}

	global.UserSrvClient = proto.NewUserClient(userConn)
}
