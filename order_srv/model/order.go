package model

import "time"

type OrderInfo struct {
	BaseModel
	User    int32  `json:"user" gorm:"type:int;index"`
	OrderSn string `gorm:"type:varchar(30);index" json:"order_sn"`
	PayType string `gorm:"type:varchar(20)" json:"pay_type"`

	Status     string    `gorm:"type:varchar(20)" json:"status"`
	TradeNo    string    `gorm:"type:varchar(100)" json:"trade_no"`
	OrderMount float32   `gorm:"type:float" json:"order_mount"`
	PayTime    time.Time `gorm:"type:datetime" json:"pay_time"`

	Address      string `json:"address" gorm:"varchar(100)"`
	SignerName   string `json:"signer_name" gorm:"varchar(20)"`
	SignerMobile string `json:"signer_mobile" gorm:"varchar(11)"`
	Post         string `json:"post" gorm:"varchar(20)"`
}

type ShoppingCart struct {
	BaseModel
	User    int32 `json:"user" gorm:"type:int;index"`
	Goods   int32 `json:"goods" gorm:"type:int;index"`
	Nums    int32 `json:"nums" gorm:"type:int"`
	Checked bool
}

type OrderGoods struct {
	BaseModel

	Order int32 `json:"order" gorm:"type:int;index"`
	Goods int32 `json:"goods" gorm:"type:int;index"`

	GoodsName  string  `json:"goods_name" gorm:"type:varchar(100);index"`
	GoodsImage string  `json:"goods_image" gorm:"type:varchar(100)"`
	GoodsPrice float32 `json:"goods_price"`
	Num        int32   `json:"num" gorm:"type:int"`
}
