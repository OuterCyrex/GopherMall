package router

import (
	"GopherMall/userop_api/api/message"
	"GopherMall/userop_api/middlewares"
	"github.com/gin-gonic/gin"
)

func InitMessageRouter(Router *gin.RouterGroup) {
	MessageRouter := Router.Group("message")
	{
		MessageRouter.GET("", middlewares.JWTAuthMiddleware(), message.List)
		MessageRouter.POST("", middlewares.JWTAuthMiddleware(), message.New)
	}
}
