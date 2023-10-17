package internal

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigKey string

const (
	API_PORT     ConfigKey = "API_PORT"
	OPEN_API_KEY ConfigKey = "OPEN_API_KEY"
)

type Config struct {
	API_PORT     string
	OPEN_API_KEY string
}

var defaultConfig = Config{
	API_PORT:     "8080",
	OPEN_API_KEY: "",
}

func InitConfig() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Could not load .env file")
	}
}

func GetConfig(key ConfigKey) string {
	value := os.Getenv(string(key))

	if value == "" {
		switch key {
		case "API_PORT":
			return defaultConfig.API_PORT
		case "OPEN_API_KEY":
			return defaultConfig.OPEN_API_KEY
		}
	}

	return value
}
