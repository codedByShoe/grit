package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppUrl     string
	DbUrl      string
	Port       string
	ViewsUrl   string
	SecretKey  string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	DbSchema   string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file does it exist?")
		return nil
	}

	return &Config{
		AppUrl:     os.Getenv("APP_URL"),
		DbUrl:      os.Getenv("DB_URL"),
		SecretKey:  os.Getenv("SECRET_KEY"),
		Port:       os.Getenv("PORT"),
		ViewsUrl:   "resources/views/app.html",
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		DbSchema:   os.Getenv("Db_Schema"),
	}
}
