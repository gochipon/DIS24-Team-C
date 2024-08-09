package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	PineconeKey   string
	OpenaiKey     string
	PineconeIndex string
}

var C *Config

func init() {
	godotenv.Load()
	config = &Config{
		DBHost:        getEnv("DB_HOST"),
		DBPort:        getEnv("DB_PORT"),
		DBUser:        getEnv("DB_USER"),
		DBPassword:    getEnv("DB_PASSWORD"),
		DBName:        getEnv("DB_NAME"),
		PineconeKey:   getEnv("PINECONE_KEY"),
		OpenaiKey:     getEnv("OPENAI_KEY"),
		PineconeIndex: getEnv("PINECONE_INDEX"),
	}
}

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic("missing env var: " + key)
	}
	return value
}
