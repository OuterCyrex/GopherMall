package model

import (
	"gorm.io/gorm"
	"time"
)

const (
	LEAVING_MESSAGES = iota + 1
	COMPLAINT
	INQUIRY
	POST_SALE
	WANT_TO_BUY
)

type BaseModel struct {
	ID        int32          `gorm:"primary key;type:int" json:"id"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"-"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"-"`
	DeleteAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
	IsDeleted bool           `gorm:"column:is_deleted" json:"-"`
}

type LeavingMessages struct {
	BaseModel

	User        int32  `gorm:"type:int;index" json:"user"`
	MessageType int32  `gorm:"type:int" json:"message_type"`
	Subject     string `gorm:"type:varchar(255)" json:"subject"`

	Message string `json:"message"`
	File    string `gorm:"type:varchar(200)" json:"file"`
}

type Address struct {
	BaseModel

	User         int32  `gorm:"type:int;index" json:"user"`
	Province     string `gorm:"type:varchar(10)" json:"province"`
	City         string `gorm:"type:varchar(10)" json:"city"`
	District     string `gorm:"type:varchar(20)" json:"district"`
	Address      string `gorm:"type:varchar(100)" json:"address"`
	SignerName   string `gorm:"type:varchar(20)" json:"signer_name"`
	SignerMobile string `gorm:"type:varchar(11)" json:"signer_mobile"`
}

type UserFav struct {
	BaseModel

	User  int32 `gorm:"type:int;index:idx_user_goods,unique" json:"user"`
	Goods int32 `gorm:"type:int;index:idx_user_goods,unique" json:"goods"`
}
