package api

import (
	"GopherMall/user_api/forms"
	"GopherMall/user_api/global"
	"GopherMall/user_api/global/response"
	"GopherMall/user_api/utils/JwtUtil"
	proto "GopherMall/user_srv/proto/.UserProto"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func removeTopStruct(fields map[string]string) map[string]string {
	resp := map[string]string{}
	for field, err := range fields {
		resp[field[strings.Index(field, ".")+1:]] = err
	}
	return resp
}

func HandleValidatorError(err error, c *gin.Context) {
	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "系统内部错误",
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
}

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

	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

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

func PasswordLogin(c *gin.Context) {
	passwordLoginForm := forms.PasswordLoginForm{}
	if err := c.ShouldBindJSON(&passwordLoginForm); err != nil {
		HandleValidatorError(err, c)
		return
	}

	if !store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, true) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	}

	userConn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d",
			global.ServerConfig.UserSrvConfig.Host,
			global.ServerConfig.UserSrvConfig.Port,
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		zap.S().Errorw(fmt.Sprintf("[PwdLogin] Connect to Grpc Server Failed"),
			"msg",
			err.Error(),
		)
	}

	userSrvClient := proto.NewUserClient(userConn)

	userInfoResp, err := userSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: passwordLoginForm.Mobile,
	})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": "电话 " + passwordLoginForm.Mobile + " 尚未注册",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "登陆失败",
				})
				zap.S().Infof("[pwdLogin] Login Failed: %v", err)
			}
		}
		return
	} else {
		if passwordValid, _ := userSrvClient.CheckPasswordInfo(context.Background(), &proto.PasswordCheck{
			Password:          passwordLoginForm.Password,
			EncryptedPassword: userInfoResp.GetPassword(),
		}); passwordValid == nil || !passwordValid.Success {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "密码错误",
			})
			return
		} else {
			token, err := JwtUtil.CreateJWT(uint(userInfoResp.Id), userInfoResp.NickName, uint(userInfoResp.Role))
			if err != nil {
				zap.S().Debugf("[PwdLogin] Generate JWT Failed: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "服务器出错",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"token": token,
				"msg":   "登陆成功",
			})
		}
	}
}
