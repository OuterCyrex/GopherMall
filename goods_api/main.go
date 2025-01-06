package main

import (
	"GopherMall/goods_api/config/policy"
	"GopherMall/goods_api/gateway/consul"
	"GopherMall/goods_api/global"
	"GopherMall/goods_api/initialize"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitSrvConnection(14, policy.RoundRobin)
	err := initialize.InitTrans("zh")
	if err != nil {
		zap.S().Panicf("init trans failed: %v", zap.Error(err))
	}

	R := initialize.Routers()

	registryId := fmt.Sprintf("%s", uuid.NewV4())

	registryClient := consul.NewRegistryClient(global.ServerConfig.Consul.Host, global.ServerConfig.Consul.Port)
	err = registryClient.Register(
		global.ServerConfig.Address,
		global.ServerConfig.Port,
		global.ServerConfig.Name,
		global.ServerConfig.Tags,
		fmt.Sprintf("%s", registryId),
	)
	if err != nil {
		zap.S().Panicf("Connect to Register Center Failed: %v", err)
	}

	zap.S().Debugf("server start... port: %d", global.ServerConfig.Port)

	go func() {
		if err := R.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
			zap.S().Panicf("server start failed : %v", err)
		}
	}()

	//终止时注销服务
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	err = registryClient.DeRegister(registryId)
	if err == nil {
		zap.S().Infof("API Gateway Deregistry Success")
	}
}
