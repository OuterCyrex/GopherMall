package initialize

import (
	"GopherMall/userop_api/global"
	AddressProto "GopherMall/userop_api/proto/.AddressProto"
	FavProto "GopherMall/userop_api/proto/.FavProto"
	goodsproto "GopherMall/userop_api/proto/.GoodsProto"
	MessageProto "GopherMall/userop_api/proto/.MessageProto"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitSrvConnection(wait uint, policy string) {
	Conn, err := grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=%ds",
			global.ServerConfig.Consul.Host,
			global.ServerConfig.Consul.Port,
			global.ServerConfig.UserOpSrv.Name,
			wait,
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"loadBalancingPolicy": "%s"}`, policy)),
	)

	if err != nil {
		zap.S().Panicf("Load Balance Init Failed: %v", err)
		return
	}

	global.AddressSrvClient = AddressProto.NewAddressClient(Conn)
	global.MessageSrvClient = MessageProto.NewMessageClient(Conn)
	global.FavSrvClient = FavProto.NewFavClient(Conn)
	global.GoodsSrvClient = goodsproto.NewGoodsClient(Conn)
}
