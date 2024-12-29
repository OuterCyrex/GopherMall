package router

import (
	"GopherMall/user_api/api"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(R *gin.RouterGroup) {
	BaseRouter := R.Group("base")
	{
		BaseRouter.GET("captcha", api.GetCaptcha)
		BaseRouter.POST("sendSms", api.SendSms)
	}
}
