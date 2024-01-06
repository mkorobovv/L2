package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mkorobovv/L2/develop/server/middleware"
	"github.com/mkorobovv/L2/develop/server/models"
	"github.com/mkorobovv/L2/develop/server/repo"
)

func UpdateEventHandler(w http.ResponseWriter, r *http.Request, c *repo.Cache) {
	if r.Method != http.MethodPost {
		middleware.ErrorLogger(w, errors.New("method not allowed"))
		return
	}

	decoder := json.NewDecoder(r.Body)
	var decoded models.Event

	if err := decoder.Decode(&decoded); err != nil {
		middleware.ErrorLogger(w, err)
		return
	}

	date := decoded.Date
	time := decoded.Time
	id := decoded.UserId
	desc := decoded.Description

	ev := middleware.NewEvent(desc, date, time, id)

	c.Update(date, time, ev)
	middleware.ResponseLogger(w, "event successfuly updated")
}
