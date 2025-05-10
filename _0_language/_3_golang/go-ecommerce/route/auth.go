package route

import (
	"github.com/gin-gonic/gin"
	"github.com/bawa-hub/go-ecommerce/controller"
)

func AuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
	}
}
