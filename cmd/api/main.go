package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/ihorlenko/weather_notifier/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/ihorlenko/weather_notifier/internal/api"
	"github.com/ihorlenko/weather_notifier/internal/config"
	"github.com/ihorlenko/weather_notifier/internal/database"
)

// @title           Weather Notifier API
// @version         1.0
// @description     API for weather notification subscription
// @host            localhost:8080
// @BasePath        /api
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

	router.GET("/ping", handlers.PingHandler)

	api := router.Group("/api")
	{
		api.GET("/weather", weatherHandler.GetWeather)
		api.POST("/subscribe", subscriptionHandler.Subscribe)
		api.GET("/confirm/:token", subscriptionHandler.Confirm)
		api.GET("/unsubscribe/:token", subscriptionHandler.Unsubscribe)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Static("/static", "./web/static")
	router.StaticFile("/", "./web/index.html")

	if err := router.Run(":" + cfg.AppConfig.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
