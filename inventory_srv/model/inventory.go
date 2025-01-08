package model

type Inventory struct {
	BaseModel
	Goods   int32 `json:"goods" gorm:"column:goods;type:int;not null;index:idx_goods"`
	Stocks  int32 `json:"stocks" gorm:"column:stocks;type:int;not null"`
	Version int32 `json:"version" gorm:"column:version;type:int;not null"`
}
