package config

import (
	"net/http"
	"os"

	"github.com/dhank77/hrai/module/users/routes"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	routes.AuthRoutes(router.Group("/auth"))
	routes.UserRoutes(router.Group("/user"))

	router.Run(":" + os.Getenv("APP_PORT"))
}
