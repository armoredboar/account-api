package controllers

import (
	"net/http"

	"github.com/eduardojonssen/account-api/server/models"
	"github.com/eduardojonssen/account-api/server/repository"
	"github.com/gin-gonic/gin"
)

func CheckAccountEndpoint(c *gin.Context) {

	email := c.Query("email")

	if email != "" {
		if exists := repository.CheckExistingEmail(email); exists == true {
			c.Status(http.StatusConflict)
			return
		}
	}

	username := c.Query("username")

	if username != "" {
		if exists := repository.CheckExistingUsername(username); exists == true {
			c.Status(http.StatusConflict)
			return
		}
	}

	c.Status(http.StatusOK)
}

func CreateAccountEndpoint(c *gin.Context) {

	var account models.Account
	if err := c.BindJSON(&account); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if alreadyRegistered := repository.CheckExistingEmail(account.Email); alreadyRegistered == true {

		c.JSON(http.StatusBadRequest, gin.H{"error": "email already registered."})

		return
	}

	// TODO: Check if email is alread registered.

	// TODO: Check if username is in use.

	// TODO: Create account.

	// TODO: Send activation email.

	c.Status(http.StatusCreated)
}
