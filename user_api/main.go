package main

import (
	"GopherMall/user_api/gateway/policy"
	"GopherMall/user_api/initialize"
	"GopherMall/user_api/utils"
	"fmt"
	"go.uber.org/zap"
	"runtime"
)

func main() {
	isDebug := runtime.GOOS == "windows"

	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitRedis()
	initialize.InitSrvConnection(14, policy.RoundRobin)
	err := initialize.InitTrans("zh")
	if err != nil {
		zap.S().Panicf("init trans failed: %v", zap.Error(err))
	}

	Port := 8080

	// 若为生产环境则使用空闲端口
	if !isDebug {
		port, err := utils.GetFreePort()
		if err != nil {
			zap.S().Panicf("get free port failed: %v", zap.Error(err))
		}
		Port = port
	}

	R := initialize.Routers()

	zap.S().Debugf("server start... port: %d", Port)

	if err := R.Run(fmt.Sprintf(":%d", Port)); err != nil {
		zap.S().Panicf("server start failed : %v", err)
	}
}
