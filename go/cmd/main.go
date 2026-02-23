package main

import (
	"database/sql"
	"shorturl-service/internal/config"
	"shorturl-service/internal/handler"
	"shorturl-service/internal/repository"
	"shorturl-service/internal/router"
	"shorturl-service/internal/service"
	"shorturl-service/internal/worker"

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

	repo := repository.NewShortUrlRepository(db)

	shortUrlService := service.NewShortUrlService(repo)
	shortUrlHandler := handler.NewShortUrlHandler(shortUrlService)

	healthService := service.NewHealthService(db)
	healthHandler := handler.NewHealthHandler(healthService)

	handler := handler.Handler{
		ShortUrlHandler: *shortUrlHandler,
		HealthHandler:   *healthHandler,
	}

	router.RegisterRoutes(&r.RouterGroup, handler)

	startWorkers(db)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}

func startWorkers(db *sql.DB) {

	workers := []worker.Worker{
		worker.NewExpiredUrlCleaner(db),
	}

	for _, w := range workers {
		go w.Process()
	}
}
