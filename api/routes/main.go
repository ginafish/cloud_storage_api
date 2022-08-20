package routes

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// Initialize server
func Run() {
	getRoutes()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Run(":5000")
}

// Initialize routes
func getRoutes() {
	v1 := router.Group("/api/v1")
	addStatusRoutes(v1)

	cloud := v1.Group("/cloud")
	addAWSRoutes(cloud)
}
