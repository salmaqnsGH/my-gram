package models

type User struct {
	GORMModel
	Username string `gorm:"not null" json:"username"`
	Email    string `gorm:"not null;uniqueIndex" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Age      uint   `gorm:"not null" json:"age"`
}
