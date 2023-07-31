package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine) {
	auth := route.Group("/auth")
	auth.GET("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
		})
	})
}
