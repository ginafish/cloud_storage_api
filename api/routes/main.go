package routes

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// Run will start the server
func Run() {
	getRoutes()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Run(":5000")
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {
	v1 := router.Group("/api/v1")
	addStatusRoutes(v1)

	cloud := v1.Group("/cloud")
	addAWSRoutes(cloud)
}
