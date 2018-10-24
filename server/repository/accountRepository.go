package repository

import "github.com/eduardojonssen/account-api/server/models"

func CheckExistingEmail(email string) bool {

	// TODO: validate against database.

	return false
}

func CheckExistingUsername(username string) bool {

	// TODO: validate against database.

	return false
}

func CreateAccount(account *models.Account) int32 {

	// TODO: generate activation code.

	// TODO: create in database.

	return 1
}

func ValidateActivationCode(email string, activationCode string) bool {

	// TODO: try to update database.

	return false
}
