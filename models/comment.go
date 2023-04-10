package models

type Comment struct {
	GORMModel
	UserID  uint   `gorm:"not null" json:"user_id"`
	PhotoID uint   `gorm:"not null" json:"photo_id"`
	Message string `gorm:"not null" json:"message"`
}

type CreateCommentInput struct {
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
	Message string `json:"message"`
}

type UpdateCommentInput struct {
	Message string `json:"message"`
}
