package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mkorobovv/L2/develop/server/middleware"
	"github.com/mkorobovv/L2/develop/server/models"
	"github.com/mkorobovv/L2/develop/server/repo"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request, c *repo.Cache) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var decoded models.Event

	if err := decoder.Decode(&decoded); err != nil {
		// TODO: make domain logger
		return
	}
	fmt.Fprintf(w, "Received event: %+v", decoded)

	desc := decoded.Description
	date := decoded.Date
	id := decoded.UserId
	tt := decoded.Time

	if _, errParse := time.Parse("2006-01-02", date); errParse != nil {
		log.Println("Error parsing date query")
		return
	}

	if _, errParse := time.Parse("15:00", tt); errParse != nil {
		log.Println("Error parsing time query")
		return
	}

	ev := middleware.NewEvent(desc, date, tt, id)
	c.Create(ev)

	log.Println("event succesfully created")
}
