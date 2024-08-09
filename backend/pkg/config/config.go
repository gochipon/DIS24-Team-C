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
	if os.Getenv("ENV_CONTENT") != "" {
		envs, err := godotenv.Unmarshal(os.Getenv("ENV_CONTENT"))
		if err != nil {
			panic(err)
		}
		C = &Config{
			DBHost:        envs["DB_HOST"],
			DBPort:        envs["DB_PORT"],
			DBUser:        envs["DB_USER"],
			DBPassword:    envs["DB_PASSWORD"],
			DBName:        envs["DB_NAME"],
			PineconeKey:   envs["PINECONE_KEY"],
			OpenaiKey:     envs["OPENAI_KEY"],
			PineconeIndex: envs["PINECONE_INDEX"],
		}
		if C.DBHost == "" || C.DBPort == "" || C.DBUser == "" || C.DBPassword == "" || C.DBName == "" || C.PineconeKey == "" || C.OpenaiKey == "" || C.PineconeIndex == "" {
			panic("missing env var")
		}
		return
	}
	C = &Config{
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
