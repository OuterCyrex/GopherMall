package initialize

import (
	"GopherMall/user_api/middlewares"
	"GopherMall/user_api/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	R := gin.Default()
	R.Use(middlewares.Cors())
	ApiGroup := R.Group("/v1")
	router.InitUserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)

	return R
}
