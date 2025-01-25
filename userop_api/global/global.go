package global

import (
	"GopherMall/userop_api/config"
	AddressProto "GopherMall/userop_api/proto/.AddressProto"
	FavProto "GopherMall/userop_api/proto/.FavProto"
	goodsproto "GopherMall/userop_api/proto/.GoodsProto"
	MessageProto "GopherMall/userop_api/proto/.MessageProto"
	ut "github.com/go-playground/universal-translator"
)

var (
	ServerConfig     config.MainConfig
	Trans            ut.Translator
	MessageSrvClient MessageProto.MessageClient
	AddressSrvClient AddressProto.AddressClient
	FavSrvClient     FavProto.FavClient
	GoodsSrvClient   goodsproto.GoodsClient
)
