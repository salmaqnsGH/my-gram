package models

import (
	"gorm.io/gorm"
)

type User struct {
	GORMModel
	Username string `gorm:"not null" json:"username"`
	Email    string `gorm:"not null;uniqueIndex" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Age      string `gorm:"not null" json:"age"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	return
}
