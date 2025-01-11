package utils

import (
	"GopherMall/inventory_srv/global"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
)

var RedisSync *redsync.Redsync

func NewRedisMutex() {
	RedisSync = redsync.New(goredis.NewPool(global.RDB))
}
