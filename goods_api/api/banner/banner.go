package banner

import (
	"GopherMall/goods_api/forms"
	"GopherMall/goods_api/global"
	proto "GopherMall/goods_api/proto/.GoodsProto"
	"GopherMall/goods_api/utils"
	"GopherMall/goods_api/validator"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func List(c *gin.Context) {
	ctx := context.Background()

	resp, err := global.GoodsSrvClient.BannerList(ctx, &proto.Empty{})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
	}

	c.JSON(http.StatusOK, resp)
}

func New(c *gin.Context) {
	ctx := context.Background()

	form := forms.BannerForm{}

	err := c.ShouldBindJSON(&form)
	if err != nil {
		validator.HandleValidatorError(err, c)
		return
	}

	resp, err := global.GoodsSrvClient.CreateBanner(ctx, &proto.BannerRequest{
		Index: form.Index,
		Image: form.Image,
		Url:   form.Url,
	})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func Delete(c *gin.Context) {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无效路径参数",
		})
		return
	}

	_, err = global.GoodsSrvClient.DeleteBanner(ctx, &proto.BannerRequest{
		Id: int32(id),
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

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无效路径参数",
		})
		return
	}

	form := forms.BannerForm{}
	err = c.ShouldBindJSON(&form)
	if err != nil {
		validator.HandleValidatorError(err, c)
		return
	}

	_, err = global.GoodsSrvClient.UpdateBanner(ctx, &proto.BannerRequest{
		Id:    int32(id),
		Image: form.Image,
		Index: form.Index,
		Url:   form.Url,
	})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}
