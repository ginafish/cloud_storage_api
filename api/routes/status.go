package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type status struct {
	app    string
	status string
}

func getAppStatus(c *gin.Context) {
	c.JSON(http.StatusOK, []status{
		{app: "api", status: "UP"},
	})
}

func addStatusRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/status")

	ping.GET("/", getAppStatus)
}
