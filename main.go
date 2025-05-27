package main

import (
	"net/http"

	function "github.com/aman-zulfiqar/MusicList-API-Task5/functions"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", function.HomeHandler)
	r.HandleFunc("/songs", function.CreateSongHandler).Methods("POST")
	r.HandleFunc("/songs", function.FetchSongsHandler).Methods("GET")
	r.HandleFunc("/songs/{id}", function.UpdateSongHandler).Methods("PUT")
	r.HandleFunc("/songs/{id}", function.DeleteSongHandler).Methods("DELETE")

	log.Info("Music Playlist API is starting on :9010")
	if err := http.ListenAndServe(":9010", r); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}