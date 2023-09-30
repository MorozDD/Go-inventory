package main

import (
	"inventory/database"
	"inventory/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	// Initialize the Gin router
	r := gin.Default()
	routes.InitializeRoutes(r)
	// Run the server
	r.Run(":8080")
}
