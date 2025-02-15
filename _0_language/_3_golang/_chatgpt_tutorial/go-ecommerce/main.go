package main

import (
	"go-ecommerce/db"
	"go-ecommerce/models"
	"go-ecommerce/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	db.ConnectDB()
	db.DB.AutoMigrate(&models.User{}) // Auto-migrate User model

	// Initialize Gin router
	r := gin.Default()

	// Register routes
	routes.UserRoutes(r)

	// Start server
	r.Run(":8090")
}
