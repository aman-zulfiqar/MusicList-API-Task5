package main

import (

	function "github.com/aman-zulfiqar/MusicList-API-Task5/functions"
	"github.com/aman-zulfiqar/MusicList-API-Task5/models"
	"github.com/aman-zulfiqar/MusicList-API-Task5/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {

	config.InitDB()

	// Auto-migrate Song model
	if err := config.DB.AutoMigrate(&models.Song{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", function.HomeHandler)
	e.POST("/songs", function.CreateSongHandler)
	e.GET("/songs", function.FetchSongsHandler)
	e.PUT("/songs/:id", function.UpdateSongHandler)
	e.DELETE("/songs/:id", function.DeleteSongHandler)

	log.Info("Music Playlist API is starting on :9010")
	if err := e.Start(":9010"); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
