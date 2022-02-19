package models

import (
	"time"
)

type Base struct {
	ID string `gorm:"primary_key:uuid;not_null" json:"id"`
	//	ID        int64 `gorm:"primary_key;auto_increment;not_null"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`
}
