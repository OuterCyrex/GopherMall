package address

import (
	"GopherMall/user_api/api"
	"GopherMall/userop_api/forms"
	"GopherMall/userop_api/global"
	"GopherMall/userop_api/middlewares"
	AddressProto "GopherMall/userop_api/proto/.AddressProto"
	"GopherMall/userop_api/validator"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func List(ctx *gin.Context) {
	request := &AddressProto.AddressRequest{}

	claims, _ := ctx.Get("claims")
	currentUser := claims.(*middlewares.CustomClaims)

	if currentUser.AuthorityId != 2 {
		userId, _ := ctx.Get("userId")
		request.UserId = int32(userId.(uint))
	}

	rsp, err := global.AddressSrvClient.GetAddressList(context.Background(), request)
	if err != nil {
		zap.S().Errorw("获取地址列表失败")
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
		reMap["province"] = value.Province
		reMap["city"] = value.City
		reMap["district"] = value.District
		reMap["address"] = value.Address
		reMap["signer_name"] = value.SignerName
		reMap["signer_mobile"] = value.SignerMobile

		result = append(result, reMap)
	}

	reMap["data"] = result

	ctx.JSON(http.StatusOK, reMap)
}

func New(ctx *gin.Context) {
	addressForm := forms.AddressForm{}
	if err := ctx.ShouldBindJSON(&addressForm); err != nil {
		validator.HandleValidatorError(err, ctx)
		return
	}

	rsp, err := global.AddressSrvClient.CreateAddress(context.Background(), &AddressProto.AddressRequest{
		Province:     addressForm.Province,
		City:         addressForm.City,
		District:     addressForm.District,
		Address:      addressForm.Address,
		SignerName:   addressForm.SignerName,
		SignerMobile: addressForm.SignerMobile,
	})

	if err != nil {
		zap.S().Errorw("新建地址失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	request := make(map[string]interface{})
	request["id"] = rsp.Id

	ctx.JSON(http.StatusOK, request)
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	_, err = global.AddressSrvClient.DeleteAddress(context.Background(), &AddressProto.AddressRequest{Id: int32(i)})
	if err != nil {
		zap.S().Errorw("删除地址失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, "")
}

func Update(ctx *gin.Context) {
	addressForm := forms.AddressForm{}
	if err := ctx.ShouldBindJSON(&addressForm); err != nil {
		validator.HandleValidatorError(err, ctx)
		return
	}

	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	_, err = global.AddressSrvClient.UpdateAddress(context.Background(), &AddressProto.AddressRequest{
		Id:           int32(i),
		Province:     addressForm.Province,
		City:         addressForm.City,
		District:     addressForm.District,
		Address:      addressForm.Address,
		SignerName:   addressForm.SignerName,
		SignerMobile: addressForm.SignerMobile,
	})
	if err != nil {
		zap.S().Errorw("更新地址失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	request := make(map[string]interface{})
	ctx.JSON(http.StatusOK, request)
}
