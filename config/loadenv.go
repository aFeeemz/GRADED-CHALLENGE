package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	fmt.Println("Loaded Environment Variables:")
	fmt.Println("MONGO_URI:", os.Getenv("MONGO_URI"))
	fmt.Println("DB_NAME:", os.Getenv("DB_NAME"))
	fmt.Println("PORT:", os.Getenv("PORT"))
}
