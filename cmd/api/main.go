package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ihorlenko/weather_notifier/internal/api"
	"github.com/ihorlenko/weather_notifier/internal/config"
	"github.com/ihorlenko/weather_notifier/internal/database"
)

func main() {

	cfg := config.LoadConfig()

	if err := database.RunMigrations(cfg); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	_, err := database.NewDBConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := gin.Default()

	api.SetupRoutes(router)

	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
