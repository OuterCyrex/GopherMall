package utils

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
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
				zap.S().Errorf("InternalServerError: %v", err)
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
				zap.S().Errorf("InternalServerError: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Message(),
				})
			}
		} else {
			zap.S().Errorf("InternalServerError: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "服务器错误",
			})
		}
	}
}
