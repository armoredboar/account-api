package controllers

import (
	"net/http"
	"strings"

	"github.com/eduardojonssen/account-api/server/contracts"
	"github.com/eduardojonssen/account-api/server/models"
	"github.com/eduardojonssen/account-api/server/repository"
	"github.com/gin-gonic/gin"
)

func CheckAccountEndpoint(c *gin.Context) {

	email := strings.TrimSpace(c.Query("email"))

	var report contracts.Report

	// Checks if the email is specified and is in use.
	if email != "" {
		if exists := repository.CheckExistingEmail(email); exists == true {
			report.AddError("100", "email", "Email already registered.")
		}
	}

	username := strings.TrimSpace(c.Query("username"))

	// Checks if the username is specified and is in use.
	if username != "" {
		if exists := repository.CheckExistingUsername(username); exists == true {
			report.AddError("101", "username", "This username has already been registered.")
		}
	}

	// If any error has occurred, exit the method with an status of conflict.
	if len(report.Errors) > 0 {
		c.JSON(http.StatusConflict, report)
		return
	}

	report.Success = true

	c.JSON(http.StatusOK, report)
}

func CreateAccountEndpoint(c *gin.Context) {

	var account models.Account
	if err := c.BindJSON(&account); err != nil {

		// TODO: Remove 'required' from model and validate each field independently.
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
