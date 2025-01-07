package forms

type BannerForm struct {
	Index int32  `json:"index" binding:"required"`
	Image string `json:"image" binding:"required,url"`
	Url   string `json:"url" binding:"required,url"`
}
