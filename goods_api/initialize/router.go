package initialize

import (
	"GopherMall/goods_api/middlewares"
	"GopherMall/goods_api/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	R := gin.Default()
	R.Use(middlewares.Cors())

	// keep-alive 检查API是否存活
	R.GET("health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "alive",
		})
	})

	ApiGroup := R.Group("/v1")
	router.InitGoodsRouter(ApiGroup)
	router.InitCategoryRouter(ApiGroup)
	router.InitBannerRouter(ApiGroup)
	router.InitBrandsRouter(ApiGroup)
	return R
}
