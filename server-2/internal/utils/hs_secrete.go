package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateHospitalSecret() (string, error) {
	bytes := make([]byte, 32) 
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}