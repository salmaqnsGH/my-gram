package models

type Photo struct {
	GORMModel
	UserID   uint   `json:"user_id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}
