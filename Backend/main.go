package main

import (
	"backend-go/config"
	"backend-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
