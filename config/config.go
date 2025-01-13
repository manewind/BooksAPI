package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type APIConfig struct {
	BooksAPIBaseURL string
}

type DBConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

type ServerConfig struct {
	ServerPort string
}

type Config struct {
	APIConfig    APIConfig
	DBConfig     DBConfig
	ServerConfig ServerConfig
}

var AppConfig *Config

func LoadConfig() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	AppConfig = &Config{
		APIConfig: APIConfig{
			BooksAPIBaseURL: getEnv("Books_API_Base_URL", "https://www.googleapis.com/books/v1/volumes"),
		},
		DBConfig: DBConfig{
			DBHost:     getEnv("DB_HOST", "db"),
			DBPort:     getEnv("DB_PORT", "5432"),
			DBUser:     getEnv("DB_USER", "aboba"),
			DBPassword: getEnv("DB_Password", "12345"),
			DBName:     getEnv("DB_NAME", "Books"),
		},
		ServerConfig: ServerConfig{
			ServerPort: getEnv("SERVER_PORT", "8080"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
