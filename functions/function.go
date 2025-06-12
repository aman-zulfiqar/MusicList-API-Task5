package functions

import (
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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

func HomeHandler(c echo.Context) error {
	log.Info("Home endpoint is hitting")
	return c.String(http.StatusOK, "Welcome to Music Playlist Api")
}

func CreateSongHandler(c echo.Context) error {
	var s Song
	if err := c.Bind(&s); err != nil {
		log.Warnf("CreateSong: invalid request %v", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	if s.Title == "" || s.Artist == "" {
		log.Warn("CreateSong: missing required fields")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Title and Artist are required"})
	}

	s.ID = uuid.New().String()
	s.Timestrap = time.Now()

	mu.Lock()
	songs[s.ID] = s
	mu.Unlock()

	log.Infof("Song Created: %s", s.Title)
	return c.JSON(http.StatusCreated, echo.Map{"message": "Song added", "song_id": s.ID})
}

func FetchSongsHandler(c echo.Context) error {
	mu.Lock()
	defer mu.Unlock()

	var playlist []Song
	for _, s := range songs {
		playlist = append(playlist, s)
	}

	log.Infof("Fetched %d songs", len(playlist))
	return c.JSON(http.StatusOK, echo.Map{"playlist": playlist})
}

func UpdateSongHandler(c echo.Context) error {
	id := c.Param("id")

	var updated Song
	if err := c.Bind(&updated); err != nil {
		log.Warnf("UpdateSong: invalid request: %v", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	mu.Lock()
	defer mu.Unlock()

	existing, ok := songs[id]
	if !ok {
		log.Warnf("UpdateSong: song not found, id=%s", id)
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Song not found"})
	}

	existing.Title = updated.Title
	existing.Artist = updated.Artist
	existing.Genre = updated.Genre
	songs[id] = existing

	log.Infof("Song updated: id=%s, title=%s", id, existing.Title)
	return c.JSON(http.StatusOK, echo.Map{"message": "Song updated"})
}

func DeleteSongHandler(c echo.Context) error {
	id := c.Param("id")

	mu.Lock()
	defer mu.Unlock()

	if _, ok := songs[id]; !ok {
		log.Warnf("DeleteSong: song not found, id=%s", id)
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Song not found"})
	}

	delete(songs, id)
	log.Infof("Song deleted: id=%s", id)
	return c.JSON(http.StatusOK, echo.Map{"message": "Song deleted"})
}
