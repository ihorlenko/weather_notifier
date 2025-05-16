package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ihorlenko/weather_notifier/internal/api"
)

func main() {
	router := gin.Default()

	api.SetupRoutes(router)

	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
