package api

import (
	"GopherMall/user_api/global"
	"GopherMall/user_api/global/response"
	proto "GopherMall/user_srv/proto/.UserProto"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"time"
)

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "Internal server error",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": e.Message(),
				})
			case codes.Unavailable:
				c.JSON(http.StatusServiceUnavailable, gin.H{
					"msg": "cannot dial rpc serve",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Message(),
				})
			}
		}
	}
}

func GetUserList(c *gin.Context) {
	userConn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d",
			global.ServerConfig.UserSrvConfig.Host,
			global.ServerConfig.UserSrvConfig.Port,
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		zap.S().Errorw(fmt.Sprintf("[GetUserList] Connect to Grpc Server Failed"),
			"msg",
			err.Error(),
		)
	}

	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	userSrvClient := proto.NewUserClient(userConn)

	rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		PageNum:  uint32(pageNum),
		PageSize: uint32(pageSize),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] Get UserList Failed")
		HandleGrpcErrorToHttp(err, c)
		return
	}

	zap.S().Debug("GetUserList start...")
	result := make([]response.UserResponse, 0)
	for _, value := range rsp.Data {
		result = append(result, response.UserResponse{
			Id:       value.GetId(),
			NickName: value.GetNickName(),
			Mobile:   value.GetMobile(),
			Gender:   value.GetGender(),
			BirthDay: response.JsonTime(time.Unix(int64(value.GetBirthDay()), 0)),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Data": result,
	})
}
