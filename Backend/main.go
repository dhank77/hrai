package main

import (
	"fmt"

	"github.com/dhank77/hrai/config"
	"github.com/dhank77/hrai/database"
)

func main() {
	database.Database()

	fmt.Println("Server running!")

	config.Router()
}
