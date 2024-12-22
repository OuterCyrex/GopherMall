package initialize

import (
	"GopherMall/user_api/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	R := gin.Default()

	ApiGroup := R.Group("/v1")
	router.InitUserRouter(ApiGroup)

	return R
}
