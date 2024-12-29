package config

type MysqlConfig struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	User     string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	DB       string `json:"db" mapstructure:"db"`
}

type ServerConfig struct {
	Name   string       `json:"name" mapstructure:"name"`
	Mysql  MysqlConfig  `json:"mysql" mapstructure:"mysql"`
	Consul ConsulConfig `json:"consul" mapstructure:"consul"`
}

type ConsulConfig struct {
	Host string `json:"host" mapstructure:"host"`
	Port int    `json:"port" mapstructure:"port"`
}
