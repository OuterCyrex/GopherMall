package initialize

import (
	"GopherMall/inventory_srv/global"
	"GopherMall/inventory_srv/utils"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			global.ServerConfig.Redis.Address,
			global.ServerConfig.Redis.Port),
		DB: global.ServerConfig.Redis.DB,
	})

	global.RDB = rdb

	if global.RDB.Ping(context.Background()).Err() != nil {
		zap.S().Panicw("redis init failed", "err", "redis init failed")
	}

	utils.NewRedisMutex()
}
