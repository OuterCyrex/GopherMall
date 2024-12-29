package router

import (
	"GopherMall/user_api/api"
	"GopherMall/user_api/middlewares"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitUserRouter(Router *gin.RouterGroup) {
	zap.S().Infof("Initialize GopherMall User Router...")
	UserRouter := Router.Group("/user")
	{
		UserRouter.GET("list", middlewares.JWTAuthMiddleware(), api.GetUserList)
		UserRouter.GET("pwdLogin", api.PasswordLogin)
	}
}
