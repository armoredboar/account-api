package repository

import (
	"strings"
)

func CheckExistingEmail(email string) bool {

	// TODO: validate against database.

	if strings.EqualFold(email, "existing@mail.com") == true {
		return true
	}

	return false
}

func CheckExistingUsername(username string) bool {

	// TODO: validate against database.

	if strings.EqualFold(username, "user") == true {
		return true
	}

	return false
}

func CreateAccount(email string, username string, hashedPassword string, activationKey string) bool {

	// TODO: create in database.

	return true
}

func GetActivationKeyStatus(activationKey string) string {

	// TODO: load activation key status.

	return "notActivated"
}

func ActivateKey(activationKey string) bool {

	// TODO: udpate database.

	return true
}
