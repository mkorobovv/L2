package middleware

import (
	time2 "time"

	"github.com/mkorobovv/L2/develop/server/models"
)

func NewEvent(des, date, time string, userId int) *models.Event {
	uid := time2.Now().Unix()
	return &models.Event{
		UserId:      userId,
		Date:        date,
		Uid:         uid,
		Description: des,
		Time:        time,
	}
}
