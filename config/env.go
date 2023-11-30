package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port       string
	Host       string
	DBHost     string
	DBPort     string
	DBDatabase string
	DBPassword string
	DBUsername string
}

func NewEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Env{
		Port:       os.Getenv("APP_PORT"),
		Host:       os.Getenv("APP_HOST"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBDatabase: os.Getenv("DB_DATABASE"),
		DBUsername: os.Getenv("DB_USERNAME"),
	}
}
