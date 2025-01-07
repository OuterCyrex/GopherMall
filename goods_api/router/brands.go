package router

import (
	"GopherMall/goods_api/api/brands"
	"github.com/gin-gonic/gin"
)

func InitBrandsRouter(Router *gin.RouterGroup) {
	BrandsRouter := Router.Group("/brands")
	{
		BrandsRouter.GET("/list", brands.BrandList)
		BrandsRouter.POST("/new", brands.NewBrand)
		BrandsRouter.DELETE("/:id", brands.DeleteBrand)
		BrandsRouter.PUT("/:id", brands.UpdateBrand)
	}

	CategoryBrandsRouter := Router.Group("/categoryBrands")
	{
		CategoryBrandsRouter.GET("/list", brands.CategoryBrandList)
		CategoryBrandsRouter.POST("/new", brands.NewCategoryBrand)
		CategoryBrandsRouter.PUT("/:id", brands.UpdateCategoryBrand)
		CategoryBrandsRouter.DELETE(":id", brands.DeleteCategoryBrand)
	}
}
