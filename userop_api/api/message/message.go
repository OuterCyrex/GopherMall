package message

import (
	"GopherMall/user_api/api"
	"GopherMall/userop_api/forms"
	"GopherMall/userop_api/global"
	MessageProto "GopherMall/userop_api/proto/.MessageProto"
	"GopherMall/userop_api/validator"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func List(ctx *gin.Context) {
	rsp, err := global.MessageSrvClient.MessageList(context.Background(), &MessageProto.MessageRequest{})
	if err != nil {
		zap.S().Errorw("获取留言失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	reMap := map[string]interface{}{
		"total": rsp.Total,
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["user_id"] = value.UserId
		reMap["type"] = value.MessageType
		reMap["subject"] = value.Subject
		reMap["message"] = value.Message
		reMap["file"] = value.File

		result = append(result, reMap)
	}
	reMap["data"] = result

	ctx.JSON(http.StatusOK, reMap)
}

func New(ctx *gin.Context) {
	messageForm := forms.MessageForm{}
	if err := ctx.ShouldBindJSON(&messageForm); err != nil {
		validator.HandleValidatorError(err, ctx)
		return
	}

	rsp, err := global.MessageSrvClient.CreateMessage(context.Background(), &MessageProto.MessageRequest{
		MessageType: messageForm.MessageType,
		Subject:     messageForm.Subject,
		Message:     messageForm.Message,
		File:        messageForm.File,
	})

	if err != nil {
		zap.S().Errorw("添加留言失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	request := make(map[string]interface{})
	request["id"] = rsp.Id

	ctx.JSON(http.StatusOK, request)
}
