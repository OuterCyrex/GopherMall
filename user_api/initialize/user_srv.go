package initialize

import (
	"GopherMall/user_api/gateway"
	"GopherMall/user_api/global"
	proto "GopherMall/user_srv/proto/.UserProto"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitSrvConnection() {

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
		zap.S().Panicf(fmt.Sprintf("Connect to Grpc Server Failed: %v"), err)
		return
	}

	global.UserSrvClient = proto.NewUserClient(userConn)
}
