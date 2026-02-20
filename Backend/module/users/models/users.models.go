package models

import "time"

type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	FirstName string    `gorm:"column:firstname" json:"first_name"`
	LastName  string    `gorm:"column:lastname" json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
