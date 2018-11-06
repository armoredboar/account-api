package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
)

// CalculateHmacSha256 converts the data string to a hashed string using SHA256 algorithm.
func CalculateHmacSha256(data string, secret string) *string {

	if data == "" || secret == "" {
		return nil
	}

	// Create a new HMAC using SHA256 algorithm.
	h := hmac.New(sha256.New, []byte(secret))

	// Write the raw data.
	h.Write([]byte(data))

	// Return the hashed string.
	result := hex.EncodeToString(h.Sum(nil))

	return &result
}

// GenerateVerificationCode creates a 6 digit random number to be used as a verification code sent by email.
func GenerateVerificationCode() string {

	return fmt.Sprintf("%03d", rand.Intn(999)) + fmt.Sprintf("%03d", rand.Intn(999))
}
