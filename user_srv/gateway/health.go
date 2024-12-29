package gateway

import (
	"GopherMall/user_srv/global"
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

func HealthCheck() {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.Consul.Host, global.ServerConfig.Consul.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		zap.S().Fatal("consul connect fail", zap.Error(err))
		return
	}

	check := &api.AgentServiceCheck{
		GRPC:                           "127.0.0.1:50051",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}

	err = client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		Name:    global.ServerConfig.Name,
		ID:      global.ServerConfig.Name,
		Port:    50051,
		Tags:    []string{"user", "grpc", "service"},
		Address: "127.0.0.1",
		Check:   check,
	})

	if err != nil {
		zap.S().Panicf("Service Register Failed: %v", err)
	}
}
