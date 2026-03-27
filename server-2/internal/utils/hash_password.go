package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const DefaultCost = bcrypt.DefaultCost 


func HashPassword(plain string) (string, error) {
	if plain == "" {
		return "", errors.New("empty password")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
func ComparePassword(hashed string, plain string) error {
	if hashed == "" || plain == "" {
		return errors.New("invalid arguments to ComparePassword")
	}
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}


func CheckPasswordHash(hashed, plain string) bool {
	return ComparePassword(hashed, plain) == nil
}