package global

import (
	"GopherMall/user_api/config"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
)

var (
	ServerConfig config.MainConfig
	RDB          redis.Client
	Trans        ut.Translator
)
