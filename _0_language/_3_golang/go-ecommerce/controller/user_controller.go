package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/bawa-hub/go-ecommerce/model"
)

func GetMe(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	u := user.(model.User)
	c.JSON(http.StatusOK, gin.H{
		"id":    u.ID,
		"name":  u.Name,
		"email": u.Email,
	})
}
