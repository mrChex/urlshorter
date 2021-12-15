package main

import (
	"log"
	"net/http"

	"github.com/mrChex/urlshorter/backend/handlers"
	"github.com/mrChex/urlshorter/backend/storage"
)

func main() {
	// config
	dbURL := "postgres://shorter:shorterpwd@localhost:5432/shorter"

	store, err := storage.NewPostgresql(dbURL)
	if err != nil {
		log.Fatalf("Error while connecting to postgresql: %v", err)
	}
	defer store.(*storage.Postgresql).Close()

	http.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlers.GetURLHandler(w, r, store)

		case "PUT":
			handlers.PutURLHandler(w, r, store)

		default:
			w.Header().Add("Allow", "GET, PUT")
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	log.Println("Starting at localhost:8080")
	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalf("Unable to start web server: %v", err)
	}

	log.Println("hello")
}
