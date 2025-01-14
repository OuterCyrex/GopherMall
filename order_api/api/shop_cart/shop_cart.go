package shop_cart

import (
	goodsproto "GopherMall/goods_api/proto/.GoodsProto"
	"GopherMall/goods_api/utils"
	invproto "GopherMall/inventory_srv/proto/.InventoryProto"
	"GopherMall/order_api/forms"
	"GopherMall/order_api/global"
	proto "GopherMall/order_api/proto/.OrderProto"
	"GopherMall/order_api/validator"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func List(c *gin.Context) {
	userId, _ := c.Get("userId")

	ctx := context.Background()

	resp, err := global.OrderSrvClient.CartItemList(ctx, &proto.UserInfo{Id: int32(userId.(uint))})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	ids := make([]int32, 0)
	for _, item := range resp.Data {
		ids = append(ids, item.GoodsId)
	}

	if len(ids) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"total": 0,
		})
		return
	}

	goodsResp, err := global.GoodsSrvClient.BatchGetGoods(ctx, &goodsproto.BatchGoodsIdInfo{Id: ids})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, goodsResp)
}

func New(c *gin.Context) {
	ctx := context.Background()

	itemForm := forms.ShopCartItemForm{}
	if err := c.ShouldBindJSON(&itemForm); err != nil {
		validator.HandleValidatorError(err, c)
		return
	}

	goods, err := global.GoodsSrvClient.GetGoodsDetail(ctx,
		&goodsproto.GoodInfoRequest{Id: itemForm.GoodsId})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	inv, err := global.InvSrvClient.InvDetail(ctx,
		&invproto.GoodsInvInfo{GoodsId: itemForm.GoodsId})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	if inv.Num < itemForm.Nums {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "商品库存不足",
		})
		return
	}

	userId, _ := c.Get("userId")
	resp, err := global.OrderSrvClient.CreateCartItem(ctx, &proto.CartItemRequest{
		UserId:     int32(userId.(uint)),
		GoodsId:    goods.Id,
		GoodsName:  goods.Name,
		GoodsImage: goods.GoodsFrontImage,
		GoodsPrice: goods.ShopPrice,
		Nums:       itemForm.Nums,
		Checked:    false,
	})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无效路径参数",
		})
		return
	}
	userId, _ := c.Get("userId")

	ctx := context.Background()

	_, err = global.OrderSrvClient.DeleteCartItem(ctx, &proto.CartItemRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(id),
	})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}

func Update(c *gin.Context) {
	ctx := context.Background()

	itemForm := forms.ShopCartItemUpdateForm{}
	err := c.ShouldBindJSON(&itemForm)
	if err != nil {
		validator.HandleValidatorError(err, c)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无效路径参数",
		})
		return
	}
	userId, _ := c.Get("userId")

	inv, err := global.InvSrvClient.InvDetail(ctx,
		&invproto.GoodsInvInfo{GoodsId: int32(id)})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	if inv.Num < itemForm.Nums {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "商品库存不足",
		})
		return
	}

	req := &proto.CartItemRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(id),
		Nums:    itemForm.Nums,
		Checked: false,
	}

	if itemForm.Checked != nil {
		req.Checked = *itemForm.Checked
	}

	_, err = global.OrderSrvClient.UpdateCartItem(ctx, req)
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}
