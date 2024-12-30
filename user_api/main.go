package main

import (
	"GopherMall/user_api/initialize"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig(true)
	initialize.InitRedis()
	initialize.InitSrvConnection()
	err := initialize.InitTrans("zh")
	if err != nil {
		zap.S().Panicf("init trans failed: %v", zap.Error(err))
	}

	R := initialize.Routers()

	Port := 8080

	zap.S().Debugf("server start... port: %d", Port)

	if err := R.Run(fmt.Sprintf(":%d", Port)); err != nil {
		zap.S().Panicf("server start failed : %v", err)
	}
}
