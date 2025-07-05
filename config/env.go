package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	if os.Getenv("GEMINI_API_KEY") == "" {
		log.Fatal("Missing GEMINI_API_KEY")
	}
}
