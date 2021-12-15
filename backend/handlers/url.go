package handlers

import (
	"encoding/json"
	"errors"
	"github.com/mrChex/urlshorter/backend/storage"
	"log"
	"net/http"
	"strings"
)

type PutURLResponse struct {
	Hash string `json:"hash"`
}

type GetURLResponse struct {
	URL string `json:"url"`
}

func cleanupAndValidateURL(url string) (string, error) {
	url = strings.ToLower(url)

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return url, errors.New("you must provide correct url")
	}

	// TODO: add more validation and cleanup rules :)

	return url, nil
}

func PutURLHandler(w http.ResponseWriter, r *http.Request, store storage.Storage) {
	r.ParseForm()
	url, err := cleanupAndValidateURL(r.FormValue("url"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	link, err := store.PutLink(url)
	if err != nil {
		log.Printf("error while store.PutLink: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	hash, err := encodeURLHash(0, link.ID)
	if err != nil {
		log.Printf("error while encoding url hash: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(PutURLResponse{Hash: hash})

	log.Println("put url", link)
}

func GetURLHandler(w http.ResponseWriter, r *http.Request, store storage.Storage) {
	_, urlID, err := decodeURLHash(r.FormValue("hash"))
	if err != nil {
		log.Printf("error while decoding url hash: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	
	link, err := store.GetLinkByID(urlID)
	if err != nil {
		if err == storage.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Printf("error while store.GetLinkByID: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GetURLResponse{URL: link.URL})
}
