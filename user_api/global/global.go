package global

import (
	"GopherMall/user_api/config"
	ut "github.com/go-playground/universal-translator"
)

var (
	ServerConfig config.MainConfig
	Trans        ut.Translator
)
