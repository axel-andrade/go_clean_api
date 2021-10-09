package models

import (
	"go_clean_api/api/shared/utils"
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        string    `gorm:"primary_key:uuid" json:"id"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = utils.GenerateUUIDV4()
	base.CreatedAt = time.Now()
	return
}
