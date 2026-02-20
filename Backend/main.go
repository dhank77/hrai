package main

import (
	"fmt"

	"github.com/dhank77/hrai/config"
	"github.com/dhank77/hrai/module/users/models"
)

func main() {
	fmt.Println("Server running!")
	// config.Router()
	DB := config.Database()

	var result models.User
	DB.Raw("SELECT id, firstname, lastname, email, password, created_at, updated_at FROM users WHERE id = ?", 1).Scan(&result)
	fmt.Println(result)
}
