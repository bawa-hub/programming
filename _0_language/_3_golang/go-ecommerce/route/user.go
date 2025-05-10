package route

import (
	"github.com/gin-gonic/gin"
	"github.com/bawa-hub/go-ecommerce/controller"
	"github.com/bawa-hub/go-ecommerce/middleware"
)

func UserRoutes(router *gin.Engine) {
	user := router.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/me", controller.GetMe)
	}
}
