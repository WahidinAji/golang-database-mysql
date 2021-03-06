package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//get env value
func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func init()  {
	//load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}