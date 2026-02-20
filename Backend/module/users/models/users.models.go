package models

import "time"

type User struct {
	ID        uint      `gorm:"primarykey" json:"id" form:"id"`
	FirstName string    `gorm:"column:first_name" json:"first_name" form:"first_name"`
	LastName  string    `gorm:"column:last_name" json:"last_name" form:"last_name"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}
