package routes

import "github.com/gin-gonic/gin"

func AuthRoutes(router *gin.RouterGroup) {
	router.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
}
