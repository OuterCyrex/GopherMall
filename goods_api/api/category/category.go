package category

import (
	"GopherMall/goods_api/forms"
	"GopherMall/goods_api/global"
	proto "GopherMall/goods_api/proto/.GoodsProto"
	"GopherMall/goods_api/utils"
	"GopherMall/goods_api/validator"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func List(c *gin.Context) {
	ctx := context.Background()

	resp, err := global.GoodsSrvClient.GetAllCategorysList(ctx, &proto.Empty{})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}
	data := make([]interface{}, 0)
	err = json.Unmarshal([]byte(resp.JsonData), &data)
	if err != nil {
		zap.S().Errorw("获取分类列表出错", "err", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务器错误",
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func Detail(c *gin.Context) {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无效路径参数",
		})
		return
	}

	resp, err := global.GoodsSrvClient.GetSubCategory(ctx, &proto.CategoryListRequest{
		Id: int32(id),
	})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func NewCategory(c *gin.Context) {
	ctx := context.Background()

	form := forms.CategoryForm{}

	err := c.ShouldBindJSON(&form)
	if err != nil {
		validator.HandleValidatorError(err, c)
		return
	}

	resp, err := global.GoodsSrvClient.CreateCategory(ctx, &proto.CategoryInfoRequest{
		Name:           form.Name,
		ParentCategory: form.ParentCategory,
		Level:          form.Level,
		IsTab:          *form.IsTab,
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

	_, err = global.GoodsSrvClient.DeleteCategory(ctx, &proto.DeleteCategoryRequest{
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

	form := forms.UpdateCategoryForm{}
	err = c.ShouldBindJSON(&form)
	if err != nil {
		validator.HandleValidatorError(err, c)
		return
	}

	_, err = global.GoodsSrvClient.UpdateCategory(ctx, &proto.CategoryInfoRequest{
		Id:    int32(id),
		Name:  form.Name,
		IsTab: *form.IsTab,
	})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})

}
