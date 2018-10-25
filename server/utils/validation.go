package utils

import (
	"fmt"
	"math/rand"
)

// GenerateVerificationCode creates a 6 digit random number to be used as a verification code sent by email.
func GenerateVerificationCode() string {

	return fmt.Sprintf("%03d", rand.Intn(999)) + fmt.Sprintf("%03d", rand.Intn(999))
}
