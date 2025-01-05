package initialize

import (
	"GopherMall/goods_api/global"
	proto "GopherMall/goods_api/proto/.GoodsProto"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// InitSrvConnection 用于实现负载均衡
func InitSrvConnection(wait uint, policy string) {
	goodsConn, err := grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=%ds",
			global.ServerConfig.Consul.Host,
			global.ServerConfig.Consul.Port,
			global.ServerConfig.GoodsSrv.Name,
			wait,
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"loadBalancingPolicy": "%s"}`, policy)),
	)

	if err != nil {
		zap.S().Panicf("Load Balance Init Failed: %v", err)
		return
	}

	global.GoodsSrvClient = proto.NewGoodsClient(goodsConn)
}
