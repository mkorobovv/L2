package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/mkorobovv/L2/develop/server/middleware"
	"github.com/mkorobovv/L2/develop/server/repo"
)

func GetEventMonthHandler(w http.ResponseWriter, r *http.Request, c *repo.Cache) {
	if r.Method != http.MethodGet {
		middleware.ErrorLogger(w, errors.New("method not allowed"))
		return
	}

	dateQuery := r.URL.Query().Get("date")

	if _, errParse := time.Parse("2006-01-02", dateQuery); errParse != nil {
		middleware.ErrorLogger(w, errParse)
		return
	}

	events := c.GetMonth(dateQuery)
	resp, err := json.MarshalIndent(events, "", "\t")
	if err != nil {
		middleware.ErrorLogger(w, err)
		return
	}
	w.Write(resp)
}
