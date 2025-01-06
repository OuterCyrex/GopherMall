package router

import (
	"GopherMall/goods_api/api/goods"
	"GopherMall/goods_api/middlewares"
	"github.com/gin-gonic/gin"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")
	{
		GoodsRouter.GET("list", goods.List)
		GoodsRouter.POST("newGoods", middlewares.JWTAuthMiddleware(), middlewares.IsAdminAuth(), goods.NewGoods)
	}
}
