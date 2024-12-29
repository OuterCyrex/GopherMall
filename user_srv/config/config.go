package config

type MysqlConfig struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	User     string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	DB       string `json:"db" mapstructure:"db"`
}

type ServerConfig struct {
	Mysql MysqlConfig `json:"mysql" mapstructure:"mysql"`
}
