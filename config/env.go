package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	return os.Getenv("MONGOURI")
}
