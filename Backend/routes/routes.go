package routes

import (
	"backend-go/controllers"
	"backend-go/middlewares"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173" // fallback ค่า default
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	auth.GET("/me", controllers.GetMe)
	auth.GET("/users", controllers.GetUsers)
	auth.GET("/users/:id", controllers.GetUserByID)
	auth.DELETE("/users/:id", controllers.DeleteUser)
	r.GET("/host", controllers.HostGame)
    r.GET("/join", controllers.JoinGame)
}
