package validator

import (
	"GopherMall/user_api/global"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

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