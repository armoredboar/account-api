package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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
