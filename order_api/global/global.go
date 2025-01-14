package global

import (
	goodsproto "GopherMall/goods_api/proto/.GoodsProto"
	invproto "GopherMall/inventory_srv/proto/.InventoryProto"
	"GopherMall/order_api/config"
	proto "GopherMall/order_api/proto/.OrderProto"
	ut "github.com/go-playground/universal-translator"
)

var (
	ServerConfig   config.MainConfig
	Trans          ut.Translator
	OrderSrvClient proto.OrderClient
	GoodsSrvClient goodsproto.GoodsClient
	InvSrvClient   invproto.InventoryClient
)
