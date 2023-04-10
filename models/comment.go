package models

type Comment struct {
	GORMModel
	UserID  uint   `gorm:"not null" json:"user_id"`
	PhotoID uint   `gorm:"not null" json:"photo_id"`
	Message string `gorm:"not null" json:"message"`
}
