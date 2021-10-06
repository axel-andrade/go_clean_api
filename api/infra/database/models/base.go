package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string    `gorm:"primary_key:uuid" json:"id"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`
	// UpdatedAt time.Time `gorm:"<-:created"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.NewV4().String()
	base.CreatedAt = time.Now()
	return
}
