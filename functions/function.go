package functions

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Song struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Artist    string    `json:"artist"`
	Genre     string    `json:"genre"`
	Timestrap time.Time `json:"timestrap"`
}

var (
	songs = make(map[string]Song)
	mu    = &sync.Mutex{}
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Home endpoint is hitting")
	w.Write([]byte("Welcome to Music Playlist Api"))
}

func CreateSongHandler(w http.ResponseWriter, r *http.Request) {
	var s Song
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		log.Warnf("CreateSong having invalid request %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if s.Title == "" || s.Artist == "" {
		log.Warn("CreateSong is missing required fields")
		http.Error(w, "Title and Artist are required", http.StatusBadRequest)
		return
	}

	s.ID = uuid.New().String()
	s.Timestrap = time.Now()

	mu.Lock()
	songs[s.ID] = s
	mu.Unlock()

	log.Infof("Song Created: %s", s.Title)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Song added", "song_id": s.ID})

}

func FetchSongsHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var playlist []Song
	for _, s := range songs {
		playlist = append(playlist, s)
	}

	log.Infof("Fetched %d songs", len(playlist))
	json.NewEncoder(w).Encode(map[string]interface{}{"playlist": playlist})
}

func UpdateSongHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var updated Song
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		log.Warnf("UpdateSong: invalid request: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	existing, ok := songs[id]
	if !ok {
		log.Warnf("UpdateSong: song not found, id=%s", id)
		http.Error(w, "Song not found", http.StatusNotFound)
		return
	}

	existing.Title = updated.Title
	existing.Artist = updated.Artist
	existing.Genre = updated.Genre
	songs[id] = existing

	log.Infof("Song updated: id=%s, title=%s", id, existing.Title)
	json.NewEncoder(w).Encode(map[string]string{"message": "Song updated"})
}

func DeleteSongHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	mu.Lock()
	defer mu.Unlock()

	if _, ok := songs[id]; !ok {
		log.Warnf("DeleteSong song is not found, id=%s", id)
		http.Error(w, "Song not found", http.StatusNotFound)
		return
	}

	delete(songs, id)
	log.Infof("Song deleted: id=%s", id)
	json.NewEncoder(w).Encode(map[string]string{"message": "Song deleted"})
}