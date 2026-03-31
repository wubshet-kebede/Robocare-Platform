package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateRandomToken creates a simple secure hex string for invitations
func GenerateRandomToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}