package routes

import (
	functions "github.com/aman-zulfiqar/MusicList-API/controllers"
	"github.com/aman-zulfiqar/MusicList-API/middleware"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Setup() *echo.Echo {
	e := echo.New()

	// Global middlewares
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())

	// Public routes
	e.GET("/", functions.HomeHandler)
	e.POST("/register", functions.RegisterUser)
	e.POST("/login", functions.Login)

	// Protected song routes
	songGroup := e.Group("/songs")
	songGroup.Use(middleware.JWTMiddleware())
	songGroup.POST("", functions.CreateSongHandler)
	songGroup.GET("", functions.FetchSongsHandler)
	songGroup.PUT("/:id", functions.UpdateSongHandler)
	songGroup.DELETE("/:id", functions.DeleteSongHandler)

	return e
}
