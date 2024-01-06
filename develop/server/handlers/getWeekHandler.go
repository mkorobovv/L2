package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mkorobovv/L2/develop/server/repo"
)

func GetEventWeekHandler(w http.ResponseWriter, r *http.Request, c *repo.Cache) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	dateQuery := r.URL.Query().Get("date")

	if _, errParse := time.Parse("2006-01-02", dateQuery); errParse != nil {
		// TODO: make logger
		return
	}

	events := c.GetWeek(dateQuery)
	resp, err := json.MarshalIndent(events, "", "\t")
	if err != nil {
		// TODO: error logger
		return
	}
	w.Write(resp)
}
