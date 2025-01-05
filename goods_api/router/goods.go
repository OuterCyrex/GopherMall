package router

import (
	"GopherMall/goods_api/api/goods"
	"github.com/gin-gonic/gin"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")
	{
		GoodsRouter.GET("list", goods.List)
	}
}
