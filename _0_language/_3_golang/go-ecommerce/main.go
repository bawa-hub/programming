package main

import (
	"github.com/bawa-hub/go-ecommerce/config"
	"github.com/bawa-hub/go-ecommerce/model"
	"github.com/bawa-hub/go-ecommerce/route"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.ConnectDatabase()

	// Run DB migrations
	config.DB.AutoMigrate(&model.User{})

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	route.AuthRoutes(router)
	route.UserRoutes(router)

	router.Run(":" + config.AppConfig.Port)
}
