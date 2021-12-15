package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mrChex/urlshorter/backend/handlers"
	"github.com/mrChex/urlshorter/backend/storage"
)

func main() {
	// config
	dbURL := os.Getenv("POSTGRES_URL")

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

	log.Println("Starting at :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Unable to start web server: %v", err)
	}

	log.Println("hello")
}
