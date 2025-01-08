package global

import (
	"GopherMall/inventory_srv/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	ServerConfig config.ServerConfig
	DB           *gorm.DB
	RDB          *redis.Client
)
