package models

type Comment struct {
	GORMModel
	UserID  uint
	PhotoID uint
	message string
}
