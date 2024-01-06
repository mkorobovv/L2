package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/mkorobovv/L2/develop/server/repo"
)

func GetEventDayHandler(w http.ResponseWriter, r *http.Request, c *repo.Cache) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	dateQuery := r.URL.Query().Get("date")

	if _, errParse := time.Parse("2006-01-02", dateQuery); errParse != nil {
		log.Println("Error parsing date query")
		return
	}

	events := c.GetDay(dateQuery)
	resp, err := json.MarshalIndent(events, "", "\t")
	if err != nil {
		log.Println("Error marshalling")
		return
	}
	w.Write(resp)
}
