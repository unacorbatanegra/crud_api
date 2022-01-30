package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var doOnce sync.Once
var config appConfig

type appConfig struct {
	AppPort string
	AppKey  string

	DBHost     string
	DBPort     string
	DBName     string
	DBUsername string
	DBPassword string
}

func init() {
	doOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println("Err: cannot load .env file")
		}
		config = load()
	})
}

func load() appConfig {
	return appConfig{
		AppPort:    os.Getenv("APP_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}
}

func Get() appConfig {
	return config
}
