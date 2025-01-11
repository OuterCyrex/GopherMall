package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int32          `gorm:"primary key;type:int" json:"id"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"-"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"-"`
	DeleteAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
	IsDeleted bool           `gorm:"column:is_deleted" json:"-"`
}
