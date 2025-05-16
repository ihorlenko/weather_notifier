package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppConfig        AppConfig
	DBConfig         DBConfig
	WeatherAPIConfig WeatherAPIConfig
}

type AppConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func (db *DBConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		db.Host, db.Port, db.User, db.Password, db.DBName, db.SSLMode)
}

type WeatherAPIConfig struct {
	APIKey string
}

func LoadConfig() *Config {
	return &Config{
		AppConfig: AppConfig{
			Port: "8080",
		},

		DBConfig: DBConfig{
			Host:     "postgres",
			Port:     "5432",
			User:     getEnv("POSTGRES_USER", "postgres"),
			Password: getEnv("POSTGRES_PASSWORD", "postgres"),
			DBName:   getEnv("POSTGRES_DB", "weather_notifier"),
			SSLMode:  "disable",
		},

		WeatherAPIConfig: WeatherAPIConfig{
			APIKey: getEnv("WEATHER_API_KEY", ""),
		},
	}
}

func getEnv(key string, defaultValue string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultValue
	}
	return env
}
