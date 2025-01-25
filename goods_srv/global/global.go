package global

import (
	"GopherMall/goods_srv/config"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

var (
	ServerConfig config.ServerConfig
	DB           *gorm.DB
	Elastic      *elastic.Client
)
