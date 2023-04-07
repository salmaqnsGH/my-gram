package models

type SocialMedia struct {
	GORMModel
	UserID         uint
	name           string
	SocialMediaUrl string
}
