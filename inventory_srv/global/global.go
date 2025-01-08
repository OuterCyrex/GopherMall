package global

import (
	"GopherMall/inventory_srv/config"
	"gorm.io/gorm"
)

var (
	ServerConfig config.ServerConfig
	DB           *gorm.DB
)
