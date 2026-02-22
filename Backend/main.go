package main

import (
	"log"

	"github.com/dhank77/hrai/config"
	"github.com/dhank77/hrai/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Database()
	config.Router()
}
