package router

import (
	"GopherMall/user_api/api"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitUserRouter(Router *gin.RouterGroup) {
	zap.S().Infof("Initialize GopherMall User Router...")
	UserRouter := Router.Group("/user")
	{
		UserRouter.GET("list", api.GetUserList)
		UserRouter.GET("pwdLogin", api.PasswordLogin)
	}
}
