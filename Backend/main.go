package main

import (
	"github.com/gin-gonic/gin"
	"backend-go/config"
	"backend-go/routes"
)

func main() {
	config.LoadEnv()
	config.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
