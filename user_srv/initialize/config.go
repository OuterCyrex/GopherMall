package initialize

import (
	"GopherMall/user_srv/global"
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig(isDebug bool) {
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
