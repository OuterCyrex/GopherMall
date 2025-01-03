package model

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int32          `gorm:"primary key;type:int" json:"id"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updatedAt"`
	DeleteAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleteAt"`
	IsDeleted bool           `gorm:"column:is_deleted" json:"isDeleted"`
}

type GormList []string

func (g *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

func (g GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}
