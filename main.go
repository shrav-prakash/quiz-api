package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shrav-prakash/quizAPI/config"
	"github.com/shrav-prakash/quizAPI/routes"
)

func main() {
	router := gin.Default()
	config.ConnectDB()
	routes.AuthRoutes(router)
	router.Run("localhost:8080")
}
