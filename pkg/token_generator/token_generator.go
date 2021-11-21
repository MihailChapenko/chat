package token_generator

import (
	"crypto/rand"
	"encoding/hex"
)

//GenerateSecureToken function return generated token
func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
