package models

type Photo struct {
	GORMModel
	UserID   uint
	title    string
	caption  string
	PhotoUrl string
}
