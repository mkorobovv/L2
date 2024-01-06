package main

import (
	"log"
	"net/http"
	"time"

	"github.com/mkorobovv/L2/develop/server"
	"github.com/mkorobovv/L2/develop/server/handlers"
	"github.com/mkorobovv/L2/develop/server/repo"
)

var port = "8000"

func main() {
	mux := http.NewServeMux()
	cache := repo.NewCache()

	mux.HandleFunc("/create_event", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateEventHandler(w, r, cache)
	})
	mux.HandleFunc("/events_for_day", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetEventDayHandler(w, r, cache)
	})
	mux.HandleFunc("/delete_event", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteEventHandler(w, r, cache)
	})
	mux.HandleFunc("/events_for_week", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetEventWeekHandler(w, r, cache)
	})
	mux.HandleFunc("/events_for_month", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetEventMonthHandler(w, r, cache)
	})

	handler := Logger(mux)

	s := server.NewServer(port, handler)
	log.Println("Starting server on port 8000")
	log.Fatal(s.Run())
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}
