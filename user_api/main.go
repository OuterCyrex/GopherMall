package main

import (
	"GopherMall/user_api/gateway/policy"
	"GopherMall/user_api/initialize"
	"fmt"
	"go.uber.org/zap"
)

func main() {

	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitRedis()
	initialize.InitSrvConnection(14, policy.RoundRobin)
	err := initialize.InitTrans("zh")
	if err != nil {
		zap.S().Panicf("init trans failed: %v", zap.Error(err))
	}

	Port := 8080

	R := initialize.Routers()

	zap.S().Debugf("server start... port: %d", Port)

	if err := R.Run(fmt.Sprintf(":%d", Port)); err != nil {
		zap.S().Panicf("server start failed : %v", err)
	}
}
