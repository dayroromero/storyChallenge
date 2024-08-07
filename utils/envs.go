package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error lading .env: %v", err)
	}
}

func GetEnvVar(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("env var %s is undefined", name)
	}
	return value
}
