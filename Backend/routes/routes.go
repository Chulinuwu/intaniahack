package routes

import (
	"backend-go/controllers"
	"backend-go/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	auth.GET("/me", controllers.GetMe)
	auth.GET("/users", controllers.GetUsers)
	auth.GET("/users/:id", controllers.GetUserByID)
	auth.DELETE("/users/:id", controllers.DeleteUser)
}
