package controllers

import (
	"net/http"
	"strings"

	"armoredboar/account-api/internal/repository"
	"armoredboar/account-api/internal/server/contracts"
	"armoredboar/account-api/internal/server/models"
	"armoredboar/account-api/internal/validation"
	"armoredboar/account-api/pkg/mail"
	"armoredboar/account-api/pkg/security"
	"armoredboar/account-api/pkg/uuid"

	"github.com/gin-gonic/gin"
)

func CheckAccountEndpoint(c *gin.Context) {

	email := strings.TrimSpace(c.Query("email"))

	var report contracts.Report

	// Checks if the email is specified and is in use.
	if email != "" {
		if repository.CheckExistingEmail(email) == true {
			report.AddError(validation.AlreadyRegistered, "email")
		}
	}

	username := strings.TrimSpace(c.Query("username"))

	// Checks if the username is specified and is in use.
	if username != "" {
		if repository.CheckExistingUsername(username) == true {
			report.AddError(validation.AlreadyRegistered, "username")
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

	var report contracts.Report

	var account models.Account
	if err := c.BindJSON(&account); err != nil {
		report.AddError(validation.CouldNotBeParsed, "post body")
		c.JSON(http.StatusBadRequest, report)
		return
	}

	// Validates the received parameters.
	if errors := account.Validate(); len(errors) > 0 {
		report.Errors = errors
		c.JSON(http.StatusBadRequest, report)
		return
	}

	// Checks for duplicated email.
	if repository.CheckExistingEmail(account.Email) == true {
		report.AddError(validation.AlreadyRegistered, "email")
	}

	// Checks for duplicated username.
	if repository.CheckExistingUsername(account.Username) == true {
		report.AddError(validation.AlreadyRegistered, "username")
	}

	// If any error has occurred, exit the method with an status of conflict.
	if len(report.Errors) > 0 {
		c.JSON(http.StatusConflict, report)
		return
	}

	// Hashes the password, so it can be stored in database.
	hashedPassword := security.CalculateHmacSha256(account.Password, "chave")

	// Create the activation key to validate the email account.
	activationKey, _ := uuid.CreateUUID()

	// Creates the account.
	success := repository.CreateAccount(account.Email, account.Username, *hashedPassword, activationKey)

	if success == false {
		report.AddError(validation.InternalServerError, "")
		c.JSON(http.StatusInternalServerError, report)
		return
	}

	// Create the activation link.
	activationURL := "http://" + c.Request.Host + "/v1/api/account/activate?key=" + activationKey

	subject := "Account activation"
	body := "<p>Hi <b>" + account.Username +
		"</b>!</p><p>Please click the link bellow to activate your account.<p/><p><a href=\"" + activationURL + "\">" + activationURL + "</a></p>"

	// Send the activation email.
	go mail.SendEmail(account.Email, subject, body)

	report.Success = true

	c.JSON(http.StatusCreated, report)
}

func ValidateActivationKeyEndpoint(c *gin.Context) {

	// Retrieves the activation key from query.
	key := strings.TrimSpace(c.Query("key"))

	var report contracts.Report

	if key == "" {
		report.AddError(validation.NullOrEmpty, "key")
		c.JSON(http.StatusBadRequest, report)
		return
	}

	statusCode := http.StatusOK

	// Get the current status of the specified activationKey.
	activationKeyStatus := repository.GetActivationKeyStatus(key)

	// If the activation key is not activated yet, performs the activation and returns.
	if activationKeyStatus == "notActivated" {

		// Activate the key async.
		go repository.ActivateKey(key)
		report.Success = true

	} else if activationKeyStatus == "activated" {

		report.AddError(validation.AlreadyActivated, "key")
		report.Success = true

	} else if activationKeyStatus == "notFound" {

		report.AddError(validation.NotFound, "key")
		statusCode = http.StatusNotFound

	} else {
		statusCode = http.StatusInternalServerError
	}

	c.JSON(statusCode, report)
}

func ResendActivationEmailEndpoint(c *gin.Context) {
	c.Status(http.StatusOK)
}

func ChangeActivationEmailEndpoint(c *gin.Context) {
	c.Status(http.StatusOK)
}
