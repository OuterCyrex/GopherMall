package router

import (
	"GopherMall/userop_api/api/address"
	"GopherMall/userop_api/middlewares"
	"github.com/gin-gonic/gin"
)

func InitAddressRouter(Router *gin.RouterGroup) {
	AddressRouter := Router.Group("address")
	{
		AddressRouter.GET("", middlewares.JWTAuthMiddleware(), address.List)
		AddressRouter.DELETE("/:id", middlewares.JWTAuthMiddleware(), address.Delete)
		AddressRouter.POST("", middlewares.JWTAuthMiddleware(), address.New)
		AddressRouter.PUT("/:id", middlewares.JWTAuthMiddleware(), address.Update)
	}
}
