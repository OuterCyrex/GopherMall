package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"net/http"
)

var store = base64Captcha.DefaultMemStore

func GetCaptcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, ans, err := cp.Generate()
	fmt.Printf("New Captcha Answer: %v\n", ans)
	if err != nil {
		zap.S().Errorw("[captcha] generate captcha fail", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成验证码出错",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"picPath":   b64s,
	})
}
