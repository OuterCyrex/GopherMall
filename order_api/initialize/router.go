package initialize

import (
	"GopherMall/order_api/middlewares"
	"GopherMall/order_api/router"
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
	router.InitOrderRouter(ApiGroup)
	router.InitShopCartRouter(ApiGroup)
	return R
}
