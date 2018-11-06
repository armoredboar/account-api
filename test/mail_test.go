package mail_test

import (
	"testing"

	"github.com/armoredboar/account-api/pkg/mail"
)

func TestValidateEmail_ValidEmail(t *testing.T) {
	if mail.ValidateEmail("example@email.com") == false {
		t.Error("Expected true")
	}
}

func TestValidateEmail_InvalidEmail(t *testing.T) {
	if mail.ValidateEmail("email with space@email") == true {
		t.Error("Expected false")
	}
}

func TestValidateEmail_EmptyEmail(t *testing.T) {
	if mail.ValidateEmail("") == true {
		t.Error("Expected false")
	}
}
