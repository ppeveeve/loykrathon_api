package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser      string
	DBPassword  string
	DBName      string
	DBHost      string
	DBPort      string
	AppPort     string
	KafkaBroker string
	KafkaTopic  string
}

func LoadConfig(env string) *Config {
	log.Printf("Loading configuration for environment: %s", env)

	if err := godotenv.Load("./config/.env." + env); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		AppPort:     os.Getenv("APP_PORT"),
		KafkaBroker: os.Getenv("KAFKA_BROKER"),
		KafkaTopic:  os.Getenv("KAFKA_TOPIC"),
	}
}
