package main

import (
	"log"
	"net/http"

	"github.com/mkorobovv/L2/develop/server"
	"github.com/mkorobovv/L2/develop/server/middleware"
)

var port = "8000"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/test/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test got"))
	})

	handler := middleware.Logger(mux)

	s := server.NewServer(port, handler)
	log.Println("Starting server on port 8000")
	log.Fatal(s.Run())
}
