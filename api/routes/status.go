package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addStatusRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/status")

	ping.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "UP")
	})
}
