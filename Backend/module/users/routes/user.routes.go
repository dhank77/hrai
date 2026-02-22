package routes

import (
	"github.com/dhank77/hrai/middleware"
	"github.com/dhank77/hrai/module/users/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	router.Use(middleware.AuthMiddleware())
	router.GET("/me", controllers.GetMe)
}
