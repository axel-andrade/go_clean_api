package models

type User struct {
	Base
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"type:varchar(255);unique" json:"email"`
	Password string `gorm:"type:varchar(255)" json:"-"`
	Token    string `gorm:"type:varchar(255);unique" json:"token"`
}
