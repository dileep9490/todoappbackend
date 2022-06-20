package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {

	if err := godotenv.Load(".env"); err != nil {
		panic("error in loading .env file")
	}
	return os.Getenv(key)
}
