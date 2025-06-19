package functions

import (
	"net/http"

	"github.com/aman-zulfiqar/MusicList-API/config"
	"github.com/aman-zulfiqar/MusicList-API/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func HomeHandler(c echo.Context) error {
	log.Info("Home endpoint hit")
	return c.String(http.StatusOK, "Welcome to Music Playlist API")
}

func CreateSongHandler(c echo.Context) error {
	var song models.Song
	if err := c.Bind(&song); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}
	if song.Title == "" || song.Artist == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Title and Artist are required"})
	}
	song.ID = uuid.New()
	query := `INSERT INTO songs (id, title, artist, genre) VALUES ($1, $2, $3, $4)`
	_, err := config.DB.Exec(query, song.ID, song.Title, song.Artist, song.Genre)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create song"})
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "Song added", "song_id": song.ID})
}

func FetchSongsHandler(c echo.Context) error {
	query := `SELECT id, title, artist, genre FROM songs`
	rows, err := config.DB.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch songs"})
	}
	defer rows.Close()

	var songs []models.Song
	for rows.Next() {
		var song models.Song
		if err := rows.Scan(&song.ID, &song.Title, &song.Artist, &song.Genre); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error scanning song"})
		}
		songs = append(songs, song)
	}

	return c.JSON(http.StatusOK, echo.Map{"playlist": songs})
}

func UpdateSongHandler(c echo.Context) error {
	id := c.Param("id")
	var input models.Song
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	var exists bool
	err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM songs WHERE id = $1)", id).Scan(&exists)
	if err != nil || !exists {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Song not found"})
	}

	query := `UPDATE songs SET title = $1, artist = $2, genre = $3 WHERE id = $4`
	_, err = config.DB.Exec(query, input.Title, input.Artist, input.Genre, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update song"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Song updated"})
}

func DeleteSongHandler(c echo.Context) error {
	id := c.Param("id")

	query := `DELETE FROM songs WHERE id = $1`
	result, err := config.DB.Exec(query, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete song"})
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not determine deletion result"})
	}
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Song not found"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Song deleted"})
}
