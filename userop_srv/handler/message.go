package handler

import (
	"GopherMall/userop_srv/global"
	"GopherMall/userop_srv/model"
	MessageProto "GopherMall/userop_srv/proto/.MessageProto"
	"context"
)

type MessageServer struct {
	MessageProto.UnimplementedMessageServer
}

func (m MessageServer) MessageList(ctx context.Context, req *MessageProto.MessageRequest) (*MessageProto.MessageListResponse, error) {
	var rsp MessageProto.MessageListResponse
	var messages []model.LeavingMessages
	var messageList []*MessageProto.MessageResponse

	result := global.DB.Where(&model.LeavingMessages{User: req.UserId}).Find(&messages)
	rsp.Total = int32(result.RowsAffected)

	for _, message := range messages {
		messageList = append(messageList, &MessageProto.MessageResponse{
			Id:          message.ID,
			UserId:      message.User,
			MessageType: message.MessageType,
			Subject:     message.Subject,
			Message:     message.Message,
			File:        message.File,
		})
	}

	rsp.Data = messageList
	return &rsp, nil
}

func (m MessageServer) CreateMessage(ctx context.Context, req *MessageProto.MessageRequest) (*MessageProto.MessageResponse, error) {
	var message model.LeavingMessages

	message.User = req.UserId
	message.MessageType = req.MessageType
	message.Subject = req.Subject
	message.Message = req.Message
	message.File = req.File

	global.DB.Save(&message)

	return &MessageProto.MessageResponse{Id: message.ID}, nil
}
