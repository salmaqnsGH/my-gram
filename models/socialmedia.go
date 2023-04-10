package models

type SocialMedia struct {
	GORMModel
	UserID         uint   `gorm:"not null" json:"user_id"`
	Name           string `gorm:"not null" json:"name"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url"`
}
