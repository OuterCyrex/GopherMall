package api

import (
	"GopherMall/user_api/forms"
	"GopherMall/user_api/global"
	"GopherMall/user_api/global/response"
	"GopherMall/user_api/utils"
	"GopherMall/user_api/utils/JwtUtil"
	proto "GopherMall/user_srv/proto/.UserProto"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
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

	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	rsp, err := global.UserSrvClient.GetUserList(context.Background(), &proto.PageInfo{
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
		utils.HandleValidatorError(err, c)
		return
	}

	if !store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, true) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	}

	userInfoResp, err := global.UserSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
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
		if passwordValid, _ := global.UserSrvClient.CheckPasswordInfo(context.Background(), &proto.PasswordCheck{
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

func Register(c *gin.Context) {
	rc := context.Background()
	registerForm := forms.RegisterForm{}
	if err := c.ShouldBindJSON(&registerForm); err != nil {
		utils.HandleValidatorError(err, c)
		return
	}

	result, err := global.RDB.Get(rc, registerForm.Mobile).Result()
	if errors.Is(err, redis.Nil) || result != registerForm.Code {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	}

	user, err := global.UserSrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		NickName: registerForm.Mobile,
		Password: registerForm.Password,
		Mobile:   registerForm.Mobile,
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":        user.Id,
		"nick_name": user.NickName,
		"msg":       "注册成功",
	})
}
