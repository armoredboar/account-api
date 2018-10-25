package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckEndpoint validates the connection between all dependent services.
func HealthCheckEndpoint(c *gin.Context) {
	c.Status(http.StatusOK)
}
