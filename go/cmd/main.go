package main

import (
	"shorturl-service/internal/config"
	"shorturl-service/internal/handler"
	"shorturl-service/internal/repository"
	"shorturl-service/internal/router"
	"shorturl-service/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	cfg := config.LoadDefaultConfig()

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	db, err := config.NewPostgresDB(cfg.DbConfig)
	if err != nil {
		panic(err)
	}

	api := r.Group("/api")

	repo := repository.NewShortUrlRepository(db)

	shortUrlService := service.NewShortUrlService(repo)
	shortUrlHandler := handler.NewShortUrlHandler(shortUrlService)

	healthService := service.NewHealthService()
	healthHandler := handler.NewHealthHandler(healthService)

	handler := handler.Handler{
		ShortUrlHandler: *shortUrlHandler,
		HealthHandler:   *healthHandler,
	}

	router.RegisterRoutes(api, handler)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}
