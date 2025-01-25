package user_fav

import (
	"GopherMall/user_api/api"
	"GopherMall/userop_api/forms"
	"GopherMall/userop_api/global"
	FavProto "GopherMall/userop_api/proto/.FavProto"
	goodsproto "GopherMall/userop_api/proto/.GoodsProto"
	"GopherMall/userop_api/validator"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func List(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	userFavRsp, err := global.FavSrvClient.GetFavList(context.Background(), &FavProto.UserFavRequest{
		UserId: int32(userId.(uint)),
	})
	if err != nil {
		zap.S().Errorw("获取收藏列表失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	ids := make([]int32, 0)
	for _, item := range userFavRsp.Data {
		ids = append(ids, item.GoodsId)
	}

	if len(ids) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"total": 0,
		})
		return
	}

	goods, err := global.GoodsSrvClient.BatchGetGoods(context.Background(), &goodsproto.BatchGoodsIdInfo{
		Id: ids,
	})
	if err != nil {
		zap.S().Errorw("[List] 批量查询【商品列表】失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	reMap := map[string]interface{}{
		"total": userFavRsp.Total,
	}

	goodsList := make([]interface{}, 0)
	for _, item := range userFavRsp.Data {
		data := gin.H{
			"id": item.GoodsId,
		}

		for _, good := range goods.Data {
			if item.GoodsId == good.Id {
				data["name"] = good.Name
				data["shop_price"] = good.ShopPrice
			}
		}

		goodsList = append(goodsList, data)
	}
	reMap["data"] = goodsList
	ctx.JSON(http.StatusOK, reMap)
}

func New(ctx *gin.Context) {
	userFavForm := forms.UserFavForm{}
	if err := ctx.ShouldBindJSON(&userFavForm); err != nil {
		validator.HandleValidatorError(err, ctx)
		return
	}

	userId, _ := ctx.Get("userId")
	_, err := global.FavSrvClient.AddUserFav(context.Background(), &FavProto.UserFavRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: userFavForm.GoodsId,
	})

	if err != nil {
		zap.S().Errorw("添加收藏记录失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	userId, _ := ctx.Get("userId")
	_, err = global.FavSrvClient.DeleteUserFav(context.Background(), &FavProto.UserFavRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(i),
	})
	if err != nil {
		zap.S().Errorw("删除收藏记录失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}

func Detail(ctx *gin.Context) {
	goodsId := ctx.Param("id")
	goodsIdInt, err := strconv.ParseInt(goodsId, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	userId, _ := ctx.Get("userId")
	_, err = global.FavSrvClient.GetUserFavDetail(context.Background(), &FavProto.UserFavRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(goodsIdInt),
	})
	if err != nil {
		zap.S().Errorw("查询收藏状态失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	ctx.Status(http.StatusOK)
}
