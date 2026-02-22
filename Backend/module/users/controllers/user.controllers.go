package controllers

import (
	"net/http"

	"github.com/dhank77/hrai/database"
	"github.com/dhank77/hrai/module/users/resources"
	"github.com/gin-gonic/gin"
)

func GetMe(c *gin.Context) {
	userId, _ := c.Get("user_id")

	var user resources.User
	database.DB.Where("id = ?", userId).First(&user)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   user,
	})
}
