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
	if err := config.DB.Create(&song).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create song"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Song added", "song_id": song.ID})
}

func FetchSongsHandler(c echo.Context) error {
	var songs []models.Song
	if err := config.DB.Find(&songs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch songs"})
	}
	return c.JSON(http.StatusOK, echo.Map{"playlist": songs})
}

func UpdateSongHandler(c echo.Context) error {
	id := c.Param("id")
	var song models.Song

	if err := config.DB.First(&song, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Song not found"})
	}

	var input models.Song
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	song.Title = input.Title
	song.Artist = input.Artist
	song.Genre = input.Genre

	if err := config.DB.Save(&song).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update song"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Song updated"})
}

func DeleteSongHandler(c echo.Context) error {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Song{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete song"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Song deleted"})
}
