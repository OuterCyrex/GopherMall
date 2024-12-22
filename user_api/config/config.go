package config

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type MainConfig struct {
	Name          string        `mapstructure:"name"`
	UserSrvConfig UserSrvConfig `mapstructure:"user_srv"`
}
