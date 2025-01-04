package model

type Category struct {
	BaseModel
	Name             string      `gorm:"type:varchar(20);not null" json:"name"`
	ParentCategoryID int32       `gorm:"type:int;not null;default:0" json:"parent_category_id"`
	ParentCategory   *Category   `json:"-"`
	SubCategory      []*Category `gorm:"foreignkey:ParentCategoryID;references:ID" json:"sub_category"`
	Level            int32       `gorm:"type:int;not null;default:1" json:"level"`
	IsTab            bool        `gorm:"default:false;not null" json:"is_tab"`
}

type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null" json:"name"`
	Logo string `gorm:"type:varchar(255);not null;default:''" json:"logo"`
}

type GoodsCategoryBrand struct {
	BaseModel
	Category   Category
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique" json:"category_id"`
	Brands     Brands
	BrandsID   int32 `gorm:"type:int;index:idx_category_brand,unique" json:"brands_id"`
}

type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(255);not null" json:"image"`
	Url   string `gorm:"type:varchar(255);not null" json:"url"`
	Index int32  `gorm:"type:int;default:1;not null"`
}

type Goods struct {
	BaseModel
	Category   Category
	CategoryID int32 `gorm:"type:int;not null" json:"category_id"`
	Brands     Brands
	BrandsID   int32 `gorm:"type:int;not null" json:"brands_id"`

	OnSale   bool `gorm:"default:false;not null" json:"on_sale"`
	ShipFree bool `gorm:"default:false;not null" json:"ship_free"`
	IsNew    bool `gorm:"default:false;not null" json:"is_new"`
	IsHot    bool `gorm:"default:false;not null" json:"is_hot"`

	Name            string   `gorm:"type:varchar(50);not null" json:"name"`
	GoodsSn         string   `gorm:"type:varchar(50);not null" json:"goods_sn"`
	ClickNum        int32    `gorm:"type:int;default:0;not null" json:"click_num"`
	SoldNum         int32    `gorm:"type:int;default:0;not null" json:"sold_num"`
	FavNum          int32    `gorm:"type:int;default:0;not null" json:"fav_num"`
	MarkPrice       float32  `gorm:"not null"`
	ShopPrice       float32  `gorm:"not null"`
	GoodsBrief      string   `gorm:"type:varchar(255);not null" json:"goods_brief"`
	Images          GormList `gorm:"type:varchar(1000);not null" json:"images"`
	DescImages      GormList `gorm:"type:varchar(1000);not null" json:"desc_images"`
	GoodsFrontImage string   `gorm:"type:varchar(255);not null" json:"goods_front_image"`
}
