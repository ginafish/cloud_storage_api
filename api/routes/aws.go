package routes

import (
	s3 "cloud-storage/pkg/aws"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addAWSRoutes(rg *gin.RouterGroup) {
	aws := rg.Group("/aws")

	aws.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, s3.ListS3("TEST"))
	})
}
