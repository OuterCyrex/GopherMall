package global

import (
	goodsproto "GopherMall/goods_srv/proto/.GoodsProto"
	invproto "GopherMall/inventory_srv/proto/.InventoryProto"
	"GopherMall/order_srv/config"
	"gorm.io/gorm"
)

var (
	ServerConfig       config.ServerConfig
	DB                 *gorm.DB
	GoodsSrvClient     goodsproto.GoodsClient
	InventorySrvClient invproto.InventoryClient
)
