package envs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Get(name string) string {

	return os.Getenv(name)
}
