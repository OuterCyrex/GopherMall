package forms

type BrandForm struct {
	Name string `json:"name" binding:"required"`
	Logo string `json:"logo" binding:"required,url"`
}

type CategoryBrandForm struct {
	CategoryId int32 `json:"category_id" binding:"required"`
	BrandId    int32 `json:"brand_id" binding:"required"`
}
