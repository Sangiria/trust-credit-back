package environment

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func GetVariable(key string) string {
	err := godotenv.Load(".env")
	
	if err != nil {
		log.Fatalf("error loading .env file")
	}

	return os.Getenv(key)
}