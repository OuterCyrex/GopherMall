package api

import (
	"GopherMall/user_api/forms"
	"GopherMall/user_api/global"
	validator2 "GopherMall/user_api/validator"
	"context"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/duke-git/lancet/v2/random"
	"github.com/duke-git/lancet/v2/validator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func SendSms(c *gin.Context) {
	rc := context.Background()

	var smsForm forms.SmsForm
	err := c.ShouldBindJSON(&smsForm)
	if err != nil {
		validator2.HandleValidatorError(err, c)
		return
	}
	if ok := validator.IsChineseMobile(smsForm.Mobile); !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "非法电话号码",
		})
		return
	}
	if global.RDB.Exists(rc, smsForm.Mobile).Val() != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码已发送",
		})
		return
	}

	client, err := dysmsapi.NewClientWithAccessKey(
		global.ServerConfig.AliyunSms.RegionId,
		global.ServerConfig.AliyunSms.AccessKeyId,
		global.ServerConfig.AliyunSms.AccessKeySecret)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		zap.S().Errorw("连接sms服务出错", err, err.Error())
		return
	}

	code := random.RandSliceFromGivenSlice[byte](
		[]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'},
		5, true,
	)

	req := requests.NewCommonRequest()
	req.Method = "POST"
	req.Scheme = "https"
	req.Domain = global.ServerConfig.AliyunSms.Domain
	req.Version = global.ServerConfig.AliyunSms.Version
	req.ApiName = "SendSms"
	req.QueryParams = map[string]string{
		"RegionId":      global.ServerConfig.AliyunSms.RegionId,
		"PhoneNumbers":  smsForm.Mobile,
		"SignName":      global.ServerConfig.AliyunSms.SignName,
		"TemplateCode":  global.ServerConfig.AliyunSms.TemplateCode,
		"TemplateParam": "{\"code\":" + string(code) + "}",
	}
	resp, err := client.ProcessCommonRequest(req)

	if err := client.DoAction(req, resp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器内部错误",
		})
		zap.S().Errorw("发送sms服务出错", err, err.Error())
		return
	}

	global.RDB.Set(rc, smsForm.Mobile, code, 5*time.Minute)

	c.JSON(http.StatusOK, gin.H{
		"msg": "短信发送成功",
	})
}
