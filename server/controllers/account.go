package controllers

import (
	"net/http"
	"strings"

	"github.com/armoredboar/account-api/server/contracts"
	"github.com/armoredboar/account-api/server/models"
	"github.com/armoredboar/account-api/server/repository"
	"github.com/armoredboar/account-api/server/utils"
	"github.com/gin-gonic/gin"
)

func CheckAccountEndpoint(c *gin.Context) {

	email := strings.TrimSpace(c.Query("email"))

	var report contracts.Report

	// Checks if the email is specified and is in use.
	if email != "" {
		if repository.CheckExistingEmail(email) == true {
			report.AddError("300", "email", "Email already registered.")
		}
	}

	username := strings.TrimSpace(c.Query("username"))

	// Checks if the username is specified and is in use.
	if username != "" {
		if repository.CheckExistingUsername(username) == true {
			report.AddError("301", "username", "Username already registered.")
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
		report.AddError("900", "body", "The post body could not be parsed.")
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
		report.AddError("300", "email", "Email already registered.")
	}

	// Checks for duplicated username.
	if repository.CheckExistingUsername(account.Username) == true {
		report.AddError("301", "username", "Username already registered.")
	}

	// If any error has occurred, exit the method with an status of conflict.
	if len(report.Errors) > 0 {
		c.JSON(http.StatusConflict, report)
		return
	}

	// Hashes the password, so it can be stored in database.
	hashedPassword := utils.CalculateHmacSha256(account.Password, "chave")

	// Create the activation key to validate the email account.
	activationKey, _ := utils.CreateUUID()

	// Creates the account.
	success := repository.CreateAccount(account.Email, account.Username, *hashedPassword, activationKey)

	if success == false {
		report.AddError("1001", "", "An internal error has ocurred. Please, try again later.")
		c.JSON(http.StatusInternalServerError, report)
		return
	}

	// Create the activation link.
	activationURL := "http://" + c.Request.Host + "/api/account/activate?key=" + activationKey

	subject := "Account activation"
	body := "<p>Hi <b>" + account.Username +
		"</b>!</p><p>Please click the link bellow to activate your account.<p/><p><a href=\"" + activationURL + "\">" + activationURL + "</a></p>"

	// Send the activation email.
	go utils.SendEmail(account.Email, subject, body)

	report.Success = true

	c.JSON(http.StatusCreated, report)
}

func ValidateActivationCodeEndpoint(c *gin.Context) {

	var report contracts.Report

	report.Success = true

	c.JSON(http.StatusOK, report)
}

func ResendActivationCodeEndpoint(c *gin.Context) {
	c.Status(http.StatusOK)
}

func ChangeActivationEmailEndpoint(c *gin.Context) {
	c.Status(http.StatusOK)
}
