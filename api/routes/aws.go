package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addAWSRoutes(rg *gin.RouterGroup) {
	aws := rg.Group("/aws")

	aws.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "UP")
	})
}
