package initialize

import (
	"GopherMall/user_api/global"
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

	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		zap.S().Panicw("Viper UnMarshal YAMLFile failed")
	}

	zap.S().Infof("Read Server Config Success: %v", global.ServerConfig)
}
