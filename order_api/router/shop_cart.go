package router

import (
	"GopherMall/order_api/api/shop_cart"
	"GopherMall/order_api/middlewares"
	"github.com/gin-gonic/gin"
)

func InitShopCartRouter(Router *gin.RouterGroup) {
	ShopCartRouter := Router.Group("shopCart").Use(middlewares.JWTAuthMiddleware())
	{
		ShopCartRouter.GET("list", shop_cart.List)
		ShopCartRouter.POST("", shop_cart.New)
		ShopCartRouter.DELETE("/:id", shop_cart.Delete)
		ShopCartRouter.PUT("/:id", shop_cart.Update)
	}
}
