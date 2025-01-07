package router

import (
	"GopherMall/goods_api/api/category"
	"github.com/gin-gonic/gin"
)

func InitCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouter := Router.Group("category")
	{
		CategoryRouter.GET("/list", category.List)
		CategoryRouter.GET("/detail/:id", category.Detail)
		CategoryRouter.POST("/new", category.NewCategory)
		CategoryRouter.DELETE("/:id", category.Delete)
		CategoryRouter.PUT("/:id", category.Update)
	}
}
