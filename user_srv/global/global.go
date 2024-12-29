package global

import (
	"GopherMall/user_srv/config"
	"gorm.io/gorm"
)

var (
	ServerConfig config.ServerConfig
	DB           *gorm.DB
)
