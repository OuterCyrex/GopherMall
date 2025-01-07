package forms

type CategoryForm struct {
	Name           string `form:"name" binding:"required"`
	ParentCategory int32  `form:"parentCategory" binding:"required"`
	Level          int32  `form:"level" binding:"required"`
	IsTab          *bool  `form:"isTab" binding:"required"`
}

type UpdateCategoryForm struct {
	Name  string `form:"name" binding:"required"`
	IsTab *bool  `form:"isTab" binding:"required"`
}
