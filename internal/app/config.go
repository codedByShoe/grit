package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type database struct {
	Url string
}

type session struct {
	ApplicationKey string
}

type server struct {
	Url  string
	Port string
	Env  string
}

type Config struct {
	Database database
	Session  session
	Server   server
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &Config{
		Database: database{
			Url: os.Getenv("DATABASE_URL"),
		},
		Session: session{
			ApplicationKey: os.Getenv("APP_KEY"),
		},
		Server: server{
			Url:  os.Getenv("URL"),
			Port: os.Getenv("PORT"),
			Env:  os.Getenv("ENVIRONMENT"),
		},
	}

	return cfg, nil
}
