package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mkorobovv/L2/develop/server/models"
	"github.com/mkorobovv/L2/develop/server/repo"
)

func DeleteEventHandler(w http.ResponseWriter, r *http.Request, c *repo.Cache) {

	if r.Method != http.MethodPost {
		// TODO: make domain logger
		return
	}

	decoder := json.NewDecoder(r.Body)
	var decoded models.Event

	if err := decoder.Decode(&decoded); err != nil {
		// TODO: make domain logger
		return
	}

	date := decoded.Date
	time := decoded.Time

	c.Delete(date, time)

	// TODO: response logger

}
