package utils

import (
	"GopherMall/inventory_srv/global"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
)

var RedisSync *redsync.Redsync

func NewRedisMutex() {
	fmt.Println(global.RDB)
	RedisSync = redsync.New(goredis.NewPool(global.RDB))
}
