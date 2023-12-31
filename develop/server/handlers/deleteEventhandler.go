package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mkorobovv/L2/develop/server/middleware"
	"github.com/mkorobovv/L2/develop/server/models"
	"github.com/mkorobovv/L2/develop/server/repo"
)

func DeleteEventHandler(w http.ResponseWriter, r *http.Request, c *repo.Cache) {

	if r.Method != http.MethodPost {
		middleware.ErrorLogger(w, errors.New("method not supported"))
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

	c.Delete(date, time)

	middleware.ResponseLogger(w, "event succesfully deleted")

}
