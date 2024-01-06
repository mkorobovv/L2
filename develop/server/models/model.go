package models

type Event struct {
	UserId      uint   `json:"user_id"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

type EventCalendar struct {
	Events []*Event `json:"events"`
}
