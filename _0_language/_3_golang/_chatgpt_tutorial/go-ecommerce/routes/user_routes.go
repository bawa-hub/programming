package routes

import (
	"go-ecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
}
