package controllers

import (
	"net/http"

	"github.com/dhank77/hrai/database"
	"github.com/dhank77/hrai/helpers"
	"github.com/dhank77/hrai/module/users/models"
	"github.com/gin-gonic/gin"
)

var user models.User

func Register(c *gin.Context) {

	c.ShouldBind(&user)
	user.Password, _ = helpers.HashPassword(user.Password)

	database.DB.Create(&user)

	if user.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	database.DB.Where("email = ?", email).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	if helpers.VerifyPassword(password, user.Password) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
	})

}
