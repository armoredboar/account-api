package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// CalculateSha256 converts the data string to a hashed string using SHA256 algorithm.
func CalculateSha256(data string, secret string) string {

	// Create a new HMAC using SHA256 algorithm.
	h := hmac.New(sha256.New, []byte(secret))

	// Write the raw data.
	h.Write([]byte(data))

	// Return the hashed string.
	return hex.EncodeToString(h.Sum(nil))
}
