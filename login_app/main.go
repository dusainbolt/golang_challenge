package main

import (
	"login-app/database"
	"login-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
