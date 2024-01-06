package models

type Event struct {
	Uid         int64  `json:"uid"`
	UserId      int    `json:"user_id"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Time        string `json:"time"`
}
