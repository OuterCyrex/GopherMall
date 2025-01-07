package router

import (
	"GopherMall/goods_api/api/goods"
	"GopherMall/goods_api/middlewares"
	"github.com/gin-gonic/gin"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")
	{
		GoodsRouter.GET("/list", goods.List)
		GoodsRouter.POST("/new", middlewares.JWTAuthMiddleware(), middlewares.IsAdminAuth(), goods.NewGoods)
		GoodsRouter.GET("/detail/:id", goods.Detail)
		GoodsRouter.DELETE("/:id", middlewares.JWTAuthMiddleware(), middlewares.IsAdminAuth(), goods.Delete)
		GoodsRouter.PATCH("/:id", middlewares.JWTAuthMiddleware(), middlewares.IsAdminAuth(), goods.UpdateStatus)
		GoodsRouter.PUT("/:id", middlewares.JWTAuthMiddleware(), middlewares.IsAdminAuth(), goods.Update)
	}
}
