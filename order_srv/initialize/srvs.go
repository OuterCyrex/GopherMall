package initialize

import (
	goodsproto "GopherMall/goods_srv/proto/.GoodsProto"
	invproto "GopherMall/inventory_srv/proto/.InventoryProto"
	"GopherMall/order_srv/global"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func initClientConn(wait uint, policy string, name string) *grpc.ClientConn {
	Conn, err := grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=%ds",
			global.ServerConfig.Consul.Host,
			global.ServerConfig.Consul.Port,
			name,
			wait,
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"loadBalancingPolicy": "%s"}`, policy)),
	)

	if err != nil {
		zap.S().Panicf("Load Balance Init Failed: %v", err)
		return nil
	}

	return Conn
}

func InitServers() {
	global.GoodsSrvClient = goodsproto.NewGoodsClient(
		initClientConn(14, "round_robin", global.ServerConfig.GoodsSrv),
	)
	global.InventorySrvClient = invproto.NewInventoryClient(
		initClientConn(14, "round_robin", global.ServerConfig.InventorySrv),
	)
}
