package main

import (
	"github.com/aman-zulfiqar/MusicList-API/config"
	"github.com/aman-zulfiqar/MusicList-API/routes"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error .env file")
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Warn("No .env file found, using system environment")
	}

	config.InitDB()

	e := routes.Setup()

	log.Info("Music Playlist API is running on http://localhost:9010")
	if err := e.Start(":9010"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
