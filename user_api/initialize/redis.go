package initialize

import (
	"GopherMall/user_api/global"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.Redis.Host, global.ServerConfig.Redis.Port),
		DB:   global.ServerConfig.Redis.DB,
	})
	global.RDB = *rdb
}
