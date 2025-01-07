package router

import (
	"GopherMall/goods_api/api/banner"
	"github.com/gin-gonic/gin"
)

func InitBannerRouter(Router *gin.RouterGroup) {
	BrandsRouter := Router.Group("/banner")
	{
		BrandsRouter.GET("/list", banner.List)
		BrandsRouter.POST("/new", banner.New)
		BrandsRouter.PUT("/:id", banner.Update)
		BrandsRouter.DELETE(":id", banner.Delete)
	}
}
