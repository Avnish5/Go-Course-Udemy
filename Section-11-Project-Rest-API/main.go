package main

import (
	"example.com/rest-api/db"

	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

// cd Section-10-Project-Rest-API
func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
