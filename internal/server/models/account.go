package models

import (
	"strings"

	"github.com/armoredboar/account-api/internal/server/contracts"
	"github.com/armoredboar/account-api/internal/validation"
	"github.com/armoredboar/account-api/pkg/mail"
)

// Account represents the data required to create a new account.
type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Validate the request parameters.
func (a *Account) Validate() []contracts.Error {

	var errors []contracts.Error

	a.Email = strings.TrimSpace(a.Email)
	a.Username = strings.TrimSpace(a.Username)
	a.Password = strings.TrimSpace(a.Password)

	// Validate the email.
	if a.Email == "" {

		errors = append(errors, contracts.CreateError(validation.NullOrEmpty, "email"))

	} else {

		if len(a.Email) > 64 {
			errors = append(errors, contracts.CreateError(validation.MaxLength, "email", 64))
		}
		if mail.ValidateEmail(a.Email) == false {
			errors = append(errors, contracts.CreateError(validation.InvalidFormat, "email"))
		}
	}

	// Validate the username.
	if a.Username == "" {

		errors = append(errors, contracts.CreateError(validation.NullOrEmpty, "username"))

	} else if len(a.Username) < 4 || len(a.Username) > 16 {

		errors = append(errors, contracts.CreateError(validation.MinAndMaxRange, "username", 4, 16))
	}

	// Validate the password.
	if a.Password == "" {

		errors = append(errors, contracts.CreateError(validation.NullOrEmpty, "password"))

	} else if len(a.Password) < 6 || len(a.Password) > 32 {

		errors = append(errors, contracts.CreateError(validation.MinAndMaxRange, "password", 6, 32))
	}

	return errors
}
