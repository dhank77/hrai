package config

import (
	"net/http"

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
	router.Run(":9000")
}
