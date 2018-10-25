package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckEndpoint(c *gin.Context) {
	c.Status(http.StatusOK)
}
