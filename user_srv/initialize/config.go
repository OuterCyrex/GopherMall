package initialize

import (
	"GopherMall/user_srv/config"
	"GopherMall/user_srv/global"
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig() {
	nacosConfig := getNacosConfig()
	sc := []constant.ServerConfig{
		{
			IpAddr: nacosConfig.Host,
			Port:   nacosConfig.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         nacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})

	if err != nil {
		zap.S().Panicf("Create Nacos Client Failed: %v", err)
		return
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: nacosConfig.DataId,
		Group:  nacosConfig.Group,
	})
	if err != nil {
		zap.S().Panicf("Get Nacos JSON Failed: %v", err)
		return
	}

	mainConfig := config.ServerConfig{}
	err = json.Unmarshal([]byte(content), &mainConfig)
	if err != nil {
		zap.S().Panicf("Unmarshal Nacos JSON Failed: %v", err)
		return
	}

	global.ServerConfig = mainConfig
}

func getNacosConfig() config.NacosConfig {
	YAMLFile := "user_srv/config/config.yaml"
	v := viper.New()
	v.SetConfigFile(YAMLFile)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panicw("Viper Read YAMLFile failed")
	}

	var nacos config.NacosConfig

	if err := v.Unmarshal(&nacos); err != nil {
		zap.S().Panicw("Viper UnMarshal YAMLFile failed")
	}

	return nacos
}

// Deprecated: Use InitConfig Instead
func InitConfig2(isDebug bool) {
	var configFilePath string
	configFilePrefix := "config"
	if isDebug {
		configFilePath = fmt.Sprintf("user_srv/config/%s-debug.yaml", configFilePrefix)
	} else {
		configFilePath = fmt.Sprintf("user_srv/config/%s-product.yaml", configFilePrefix)
	}
	v := viper.New()
	v.SetConfigFile(configFilePath)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error Reading Config File: %s \n", err))
	}

	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(fmt.Errorf("Fatal error Unmarshal Config File: %s \n", err))
	}
}
