package global

import (
	"GopherMall/goods_api/config"
	proto "GopherMall/goods_api/proto/.GoodsProto"
	ut "github.com/go-playground/universal-translator"
)

var (
	ServerConfig   config.MainConfig
	Trans          ut.Translator
	GoodsSrvClient proto.GoodsClient
)
