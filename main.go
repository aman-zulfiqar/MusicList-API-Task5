package main

import (
	"github.com/aman-zulfiqar/MusicList-API/config"
	functions "github.com/aman-zulfiqar/MusicList-API/functions"
	"github.com/aman-zulfiqar/MusicList-API/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {
	config.InitDB()
	
	if err := config.DB.AutoMigrate(&models.Song{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", functions.HomeHandler)
	e.POST("/songs", functions.CreateSongHandler)
	e.GET("/songs", functions.FetchSongsHandler)
	e.PUT("/songs/:id", functions.UpdateSongHandler)
	e.DELETE("/songs/:id", functions.DeleteSongHandler)

	log.Info("Music Playlist API is starting on :9010")
	if err := e.Start(":9010"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
