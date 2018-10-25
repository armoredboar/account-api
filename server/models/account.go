package models

import (
	"strings"

	"github.com/eduardojonssen/account-api/server/contracts"
	"github.com/eduardojonssen/account-api/server/utils"
)

// Account represents the data required to create a new account.
type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Validate the received account information.
func (a *Account) Validate() []contracts.Error {

	var errors []contracts.Error

	a.Email = strings.TrimSpace(a.Email)
	a.Username = strings.TrimSpace(a.Username)
	a.Password = strings.TrimSpace(a.Password)

	// Validates the email.
	if a.Email == "" || len(a.Email) > 64 {
		errors = append(errors, contracts.Error{Code: "100", Field: "email", Message: "The email cannot be null and may not exceed 64 characters."})
	} else if utils.ValidateEmail(a.Email) == false {
		errors = append(errors, contracts.Error{Code: "200", Field: "email", Message: "The specified email is not in a valid format."})
	}

	// Validates the username.
	if len(a.Username) < 4 || len(a.Username) > 16 {
		errors = append(errors, contracts.Error{Code: "201", Field: "username", Message: "Username must have between 4 and 16 characters."})
	}

	// Validates the password.
	if len(a.Password) < 6 || len(a.Password) > 32 {
		errors = append(errors, contracts.Error{Code: "202", Field: "password", Message: "Password must have between 6 and 32 characters."})
	}

	return errors
}
