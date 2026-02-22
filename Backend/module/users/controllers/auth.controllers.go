package controllers

import (
	"net/http"

	"github.com/dhank77/hrai/database"
	"github.com/dhank77/hrai/helpers"
	"github.com/dhank77/hrai/module/users/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid request, " + err.Error(),
		})
		return
	}

	hashedPassword, _ := helpers.HashPassword(user.Password)
	user.Password = hashedPassword

	database.DB.Create(&user)

	if user.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "User created successfully",
	})
}

func Login(c *gin.Context) {
	var user models.User
	email := c.PostForm("email")
	password := c.PostForm("password")

	database.DB.Where("email = ?", email).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "User not found",
		})
		return
	}

	if err := helpers.VerifyPassword(password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": "Invalid password",
		})
		return
	}

	user.Password = "[PASSWORD]"
	token, _ := helpers.GenerateJWT(user)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"user":   user,
		"token":  "Bearer " + token,
	})
}
