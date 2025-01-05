package config

// 统一命名：
// 所有结构体在定义时需要加Config后缀，如配置Redis的结构体需要写为RedisConfig
// MainConfig是所有配置文件的总体，其内部的字段无需加Config
// yaml内对应的key即是字段的小写，如Host字段在yaml里的key就是host
// 多单词字段 需要将大驼峰命名法改为小驼峰命名法，即首字母改为小写
// Viper在使用蛇形命名法时存在一些问题，故不使用蛇形

type MainConfig struct {
	Name     string         `mapstructure:"name" json:"name"`
	Address  string         `mapstructure:"address" json:"address"`
	Port     int            `mapstructure:"port" json:"port"`
	Tags     []string       `mapstructure:"tags" json:"tags"`
	JwtKey   string         `mapstructure:"jwtKey" json:"jwtKey"`
	Consul   ConsulConfig   `mapstructure:"consul" json:"consul"`
	GoodsSrv GoodsSrvConfig `mapstructure:"goodsSrv" json:"goodsSrv"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type NacosConfig struct {
	Host      string `mapstruct:"host"`
	Port      uint64 `mapstruct:"port"`
	Namespace string `mapstruct:"namespace"`
	DataId    string `mapstruct:"dataId"`
	Group     string `mapstruct:"group"`
}

type GoodsSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
