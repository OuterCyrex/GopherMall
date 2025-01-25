package router

import (
	fav "GopherMall/userop_api/api/fav"
	"GopherMall/userop_api/middlewares"
	"github.com/gin-gonic/gin"
)

func InitUserFavRouter(Router *gin.RouterGroup) {
	UserFavRouter := Router.Group("fav")
	{
		UserFavRouter.DELETE("/:id", middlewares.JWTAuthMiddleware(), fav.Delete)
		UserFavRouter.GET("/:id", middlewares.JWTAuthMiddleware(), fav.Detail)
		UserFavRouter.POST("", middlewares.JWTAuthMiddleware(), fav.New)
		UserFavRouter.GET("", middlewares.JWTAuthMiddleware(), fav.List)
	}
}
