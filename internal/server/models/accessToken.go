package models

import (
	"github.com/armoredboar/account-api/internal/server/contracts"
	"github.com/armoredboar/account-api/internal/validation"
)

type AccessToken struct {
	GrantType    string `form:"grant_type" json:"grant_type"`
	ClientID     string `form:"client_id" json:"client_id"`
	ClientSecret string `form:"client_secret" json:"client_secret"`
	Scope        string `form:"scope" json:"scope"`
	Username     string `form:"username" json:"username"`
	Password     string `form:"password" json:"password"`
}

// Validate the received request parameters.
func (a *AccessToken) Validate() []contracts.Error {

	var errors []contracts.Error

	// Validate the GrantType
	if a.GrantType == "" {
		errors = append(errors, contracts.CreateError(validation.NullOrEmpty, "grant_type"))
	}

	// Validate the ClientID
	if a.ClientID == "" {
		errors = append(errors, contracts.CreateError(validation.NullOrEmpty, "client_id"))
	}

	// Validate the ClientSecret
	if a.ClientSecret == "" {
		errors = append(errors, contracts.CreateError(validation.NullOrEmpty, "client_secret"))
	}

	// Validate the Username
	if a.Username == "" {
		errors = append(errors, contracts.CreateError(validation.NullOrEmpty, "username"))
	}

	// Validate the Password
	if a.Password == "" {
		errors = append(errors, contracts.CreateError(validation.NullOrEmpty, "password"))
	}

	return errors
}
