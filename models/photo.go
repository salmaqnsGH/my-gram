package models

type Photo struct {
	GORMModel
	UserID   uint   `form:"user_id"`
	Title    string `form:"title"`
	Caption  string `form:"caption"`
	PhotoUrl string `form:"photo_url"`
}
