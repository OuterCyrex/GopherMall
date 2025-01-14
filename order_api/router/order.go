package router

import (
	"GopherMall/order_api/api/order"
	"GopherMall/order_api/middlewares"
	"github.com/gin-gonic/gin"
)

func InitOrderRouter(Router *gin.RouterGroup) {
	OrderRouter := Router.Group("order").Use(middlewares.JWTAuthMiddleware())
	{
		OrderRouter.GET("/list", order.List)
		OrderRouter.GET("/:id", order.Detail)
		OrderRouter.POST("", order.New)
	}
}
