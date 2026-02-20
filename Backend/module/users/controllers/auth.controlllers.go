package controllers

import (
	"github.com/dhank77/hrai/database"
	"github.com/dhank77/hrai/module/users/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	c.ShouldBind(&user)
	database.DB.Create(&user)

	if user.ID == 0 {
		c.JSON(500, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	user.Password = ""

	c.JSON(200, gin.H{
		"message": "User created successfully",
	})
}
