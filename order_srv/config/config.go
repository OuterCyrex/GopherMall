package config

type MysqlConfig struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	User     string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	DB       string `json:"db" mapstructure:"db"`
}

type ServerConfig struct {
	Name         string       `json:"name" mapstructure:"name"`
	Addr         string       `json:"addr" mapstructure:"addr"`
	Tags         []string     `json:"tags" mapstructure:"tags"`
	Mysql        MysqlConfig  `json:"mysql" mapstructure:"mysql"`
	Consul       ConsulConfig `json:"consul" mapstructure:"consul"`
	GoodsSrv     string       `json:"goods_srv" mapstructure:"goods_srv"`
	InventorySrv string       `json:"inventory_srv" mapstructure:"inventory_srv"`
}

type ConsulConfig struct {
	Host string `json:"host" mapstructure:"host"`
	Port int    `json:"port" mapstructure:"port"`
}

type NacosConfig struct {
	Host      string `mapstruct:"host"`
	Port      uint64 `mapstruct:"port"`
	Namespace string `mapstruct:"namespace"`
	DataId    string `mapstruct:"dataId"`
	Group     string `mapstruct:"group"`
}
