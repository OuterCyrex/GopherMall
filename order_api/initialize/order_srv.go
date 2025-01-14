package initialize

import (
	goodsproto "GopherMall/goods_api/proto/.GoodsProto"
	invproto "GopherMall/inventory_srv/proto/.InventoryProto"
	"GopherMall/order_api/global"
	proto "GopherMall/order_api/proto/.OrderProto"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// InitSrvConnection 用于实现负载均衡
func InitSrvConnection(wait uint, policy string) {
	global.OrderSrvClient = proto.NewOrderClient(newClient(global.ServerConfig.OrderSrv.Name, wait, policy))
	global.GoodsSrvClient = goodsproto.NewGoodsClient(newClient(global.ServerConfig.GoodsSrv.Name, wait, policy))
	global.InvSrvClient = invproto.NewInventoryClient(newClient(global.ServerConfig.InvSrv.Name, wait, policy))
}

func newClient(srvName string, wait uint, policy string) *grpc.ClientConn {
	Conn, err := grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=%ds",
			global.ServerConfig.Consul.Host,
			global.ServerConfig.Consul.Port,
			srvName,
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
