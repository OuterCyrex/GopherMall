package initialize

import (
	"GopherMall/user_api/config"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig(isDebug bool) {
	YAMLPrefix := "config"
	YAMLFile := fmt.Sprintf("user_api/config/%s-product.yaml", YAMLPrefix)
	if isDebug {
		YAMLFile = fmt.Sprintf("user_api/config/%s-debug.yaml", YAMLPrefix)
	}

	v := viper.New()
	v.SetConfigFile(YAMLFile)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panicw("Viper Read YAMLFile failed")
	}

	serverConfig := config.MainConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		zap.S().Panicw("Viper UnMarshal YAMLFile failed")
	}
}
