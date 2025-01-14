package order

import (
	"GopherMall/order_api/forms"
	"GopherMall/order_api/global"
	"GopherMall/order_api/middlewares"
	proto "GopherMall/order_api/proto/.OrderProto"
	"GopherMall/order_api/utils"
	"GopherMall/order_api/validator"
	"context"
	lancet "github.com/duke-git/lancet/v2/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func List(c *gin.Context) {
	claim, _ := c.Get("claims")
	userInfo := claim.(*middlewares.CustomClaims)
	ctx := context.Background()

	req := proto.OrderFilterRequest{}

	if userInfo.AuthorityId == 1 {
		req.UserId = int32(userInfo.ID)
	}

	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	req.Pages = int32(pageNum)
	req.PagePerNums = int32(pageSize)

	resp, err := global.OrderSrvClient.OrderList(ctx, &req)
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func New(c *gin.Context) {
	ctx := context.Background()
	orderForm := forms.CreateOrderForm{}
	if err := c.ShouldBindJSON(&orderForm); err != nil {
		validator.HandleValidatorError(err, c)
		return
	}

	if !lancet.IsChineseMobile(orderForm.Mobile) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无效的电话号码",
		})
		return
	}
	userId, _ := c.Get("userId")

	resp, err := global.OrderSrvClient.CreateOrder(ctx, &proto.OrderRequest{
		UserId:  int32(userId.(uint)),
		Address: orderForm.Address,
		Name:    orderForm.Name,
		Mobile:  orderForm.Mobile,
		Post:    orderForm.Post,
	})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func Detail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无效路径参数",
		})
		return
	}

	req := &proto.OrderRequest{}
	claim, _ := c.Get("claims")
	userInfo := claim.(*middlewares.CustomClaims)
	ctx := context.Background()

	if userInfo.AuthorityId == 1 {
		req.UserId = int32(userInfo.ID)
	}

	req.Id = int32(id)

	resp, err := global.OrderSrvClient.OrderDetail(ctx, req)
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, resp)
}
