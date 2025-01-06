package goods

import (
	"GopherMall/goods_api/forms"
	"GopherMall/goods_api/global"
	proto "GopherMall/goods_api/proto/.GoodsProto"
	"GopherMall/goods_api/validator"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
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

func List(c *gin.Context) {
	minPrice, _ := strconv.Atoi(c.DefaultQuery("minPrice", "0"))
	maxPrice, _ := strconv.Atoi(c.DefaultQuery("maxPrice", "1000000"))
	categoryId, _ := strconv.Atoi(c.DefaultQuery("categoryId", "0"))
	brandId, _ := strconv.Atoi(c.DefaultQuery("brandId", "0"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	req := &proto.GoodsFilterRequest{
		PriceMin:    int32(minPrice),
		PriceMax:    int32(maxPrice),
		TopCategory: int32(categoryId),
		Pages:       int32(pageNum),
		PagePerNums: int32(pageSize),
		KeyWords:    c.DefaultQuery("keyWords", ""),
		Brand:       int32(brandId),
		IsHot:       false,
		IsNew:       false,
		IsTab:       false,
	}

	if q := c.DefaultQuery("isHot", "0"); q == "1" {
		req.IsHot = true
	}
	if q := c.DefaultQuery("isNew", "0"); q == "1" {
		req.IsNew = true
	}
	if q := c.DefaultQuery("isTab", "0"); q == "1" {
		req.IsTab = true
	}

	resp, err := global.GoodsSrvClient.GoodsList(context.Background(), req)
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": resp.Total,
		"data":  resp.Data,
	})
}

func NewGoods(c *gin.Context) {
	goodsForm := forms.GoodsForm{}
	if err := c.ShouldBindJSON(&goodsForm); err != nil {
		validator.HandleValidatorError(err, c)
		return
	}

	ctx := context.Background()
	resp, err := global.GoodsSrvClient.CreateGoods(ctx, &proto.CreateGoodsInfo{
		Name:            goodsForm.Name,
		GoodsSn:         goodsForm.GoodsSn,
		Stocks:          goodsForm.Stocks,
		MarketPrice:     goodsForm.MarketPrice,
		ShopPrice:       goodsForm.ShopPrice,
		GoodsBrief:      goodsForm.GoodsBrief,
		ShipFree:        *goodsForm.ShipFree,
		Images:          goodsForm.Images,
		DescImages:      goodsForm.DescImages,
		GoodsFrontImage: goodsForm.FrontImage,
		CategoryId:      goodsForm.CategoryId,
		BrandId:         goodsForm.BrandId,
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func Detail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无效参数",
		})
		return
	}

	ctx := context.Background()

	r, err := global.GoodsSrvClient.GetGoodsDetail(ctx, &proto.GoodInfoRequest{
		Id: int32(id),
	})

	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": r,
	})
}

func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无效参数",
		})
		return
	}

	ctx := context.Background()

	_, err = global.GoodsSrvClient.DeleteGoods(ctx, &proto.DeleteGoodsInfo{
		Id: int32(id),
	})

	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}
