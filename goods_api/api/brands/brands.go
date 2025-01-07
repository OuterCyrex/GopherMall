package brands

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

func BrandList(c *gin.Context) {
	ctx := context.Background()

	pages, err := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pagePerNums, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无效路径参数",
		})
		return
	}

	resp, err := global.GoodsSrvClient.BrandList(ctx, &proto.BrandFilterRequest{
		Pages:       int32(pages),
		PagePerNums: int32(pagePerNums),
	})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func NewBrand(c *gin.Context) {
	ctx := context.Background()

	form := forms.BrandForm{}

	err := c.ShouldBindJSON(&form)
	if err != nil {
		validator.HandleValidatorError(err, c)
		return
	}

	resp, err := global.GoodsSrvClient.CreateBrand(ctx, &proto.BrandRequest{
		Name: form.Name,
		Logo: form.Logo,
	})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func DeleteBrand(c *gin.Context) {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无效路径参数",
		})
		return
	}

	_, err = global.GoodsSrvClient.DeleteBrand(ctx, &proto.BrandRequest{
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

func UpdateBrand(c *gin.Context) {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无效路径参数",
		})
		return
	}

	form := forms.BrandForm{}
	err = c.ShouldBindJSON(&form)
	if err != nil {
		validator.HandleValidatorError(err, c)
		return
	}

	_, err = global.GoodsSrvClient.UpdateBrand(ctx, &proto.BrandRequest{
		Id:   int32(id),
		Name: form.Name,
		Logo: form.Logo,
	})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}

func CategoryBrandList(c *gin.Context) {
	ctx := context.Background()

	pages, err := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pagePerNums, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无效路径参数",
		})
		return
	}

	resp, err := global.GoodsSrvClient.CategoryBrandList(ctx, &proto.CategoryBrandFilterRequest{
		Pages:       int32(pages),
		PagePerNums: int32(pagePerNums),
	})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func NewCategoryBrand(c *gin.Context) {
	ctx := context.Background()

	form := forms.CategoryBrandForm{}

	err := c.ShouldBindJSON(&form)
	if err != nil {
		validator.HandleValidatorError(err, c)
		return
	}

	resp, err := global.GoodsSrvClient.CreateCategoryBrand(ctx, &proto.CategoryBrandRequest{
		CategoryId: form.CategoryId,
		BrandId:    form.BrandId,
	})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func DeleteCategoryBrand(c *gin.Context) {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无效路径参数",
		})
		return
	}

	_, err = global.GoodsSrvClient.DeleteCategoryBrand(ctx, &proto.CategoryBrandRequest{
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

func UpdateCategoryBrand(c *gin.Context) {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无效路径参数",
		})
		return
	}

	form := forms.CategoryBrandForm{}
	err = c.ShouldBindJSON(&form)
	if err != nil {
		validator.HandleValidatorError(err, c)
		return
	}

	_, err = global.GoodsSrvClient.UpdateCategoryBrand(ctx, &proto.CategoryBrandRequest{
		Id:         int32(id),
		CategoryId: form.CategoryId,
		BrandId:    form.BrandId,
	})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}
