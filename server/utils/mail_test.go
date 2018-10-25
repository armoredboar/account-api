package utils_test

import (
	"testing"

	"github.com/armoredboar/account-api/server/utils"
)

func TestValidateEmail_ValidEmail(t *testing.T) {
	if utils.ValidateEmail("example@email.com") == false {
		t.Error("Expected true")
	}
}

func TestValidateEmail_InvalidEmail(t *testing.T) {
	if utils.ValidateEmail("email with space@email") == true {
		t.Error("Expected false")
	}
}

func TestValidateEmail_EmptyEmail(t *testing.T) {
	if utils.ValidateEmail("") == true {
		t.Error("Expected false")
	}
}
