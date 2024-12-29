package config

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type MainConfig struct {
	Name          string        `mapstructure:"name"`
	UserSrvConfig UserSrvConfig `mapstructure:"user_srv"`
	JwtKey        string        `mapstructure:"jwt_key"`
	AliyunSms     AliyunSms     `mapstructure:"aliyun_sms"`
	Redis         Redis         `mapstructure:"redis"`
}

type Redis struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	DB   int    `mapstructure:"db"`
}

type AliyunSms struct {
	RegionId        string `mapstruct:"regionId"`
	AccessKeyId     string `mapstruct:"accessKeyId"`
	AccessKeySecret string `mapstruct:"accessKeySecret"`
	Domain          string `mapstruct:"Domain"`
	Version         string `mapstruct:"Version"`
	SignName        string `mapstruct:"SignName"`
	TemplateCode    string `mapstruct:"TemplateCode"`
}
