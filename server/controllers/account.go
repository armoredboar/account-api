package controllers

import (
	"net/http"

	"github.com/eduardojonssen/account-api/server/models"
	"github.com/gin-gonic/gin"
)

func CreateAccountEndpoint(c *gin.Context) {

	var account models.Account
	if err := c.BindJSON(&account); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	// TODO: Check if email is alread registered.

	// TODO: Check if username is in use.

	// TODO: Create account.

	// TODO: Send activation email.

	c.Status(http.StatusCreated)
}
