package config

// 统一命名：
// 所有结构体在定义时需要加Config后缀，如配置Redis的结构体需要写为RedisConfig
// MainConfig是所有配置文件的总体，其内部的字段无需加Config
// yaml内对应的key即是字段的小写，如Host字段在yaml里的key就是host
// 多单词字段 需要将大驼峰命名法改为小驼峰命名法，即首字母改为小写
// Viper在使用蛇形命名法时存在一些问题，故不使用蛇形

type MainConfig struct {
	Name      string          `mapstructure:"name"`
	UserSrv   UserSrvConfig   `mapstructure:"userSrv"`
	JwtKey    string          `mapstructure:"jwtKey"`
	AliyunSms AliyunSmsConfig `mapstructure:"aliyunSms"`
	Redis     RedisConfig     `mapstructure:"redis"`
	Consul    ConsulConfig    `mapstructure:"consul"`
}

type UserSrvConfig struct {
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	DB   int    `mapstructure:"db"`
}

type AliyunSmsConfig struct {
	RegionId        string `mapstruct:"regionId"`
	AccessKeyId     string `mapstruct:"accessKeyId"`
	AccessKeySecret string `mapstruct:"accessKeySecret"`
	Domain          string `mapstruct:"domain"`
	Version         string `mapstruct:"version"`
	SignName        string `mapstruct:"signName"`
	TemplateCode    string `mapstruct:"templateCode"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
