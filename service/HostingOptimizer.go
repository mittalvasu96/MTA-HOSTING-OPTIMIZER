package service

import (
	"mta-hosting/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMTAHost(c *gin.Context) {
	var response helper.HostingResponse
	response.Hostname, _ = helper.FetchActiveMTA()
	response.Code = "200"
	c.JSON(http.StatusOK, response)
}

func HealthCheck(c *gin.Context) {
	c.JSON(200, "Okay!")
}
