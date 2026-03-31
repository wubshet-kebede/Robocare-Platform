package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func jwtSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func GenerateJWT(ID string, Hospital_id string, role string, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"id":   ID,
		"hospital_id": Hospital_id,
		"role": role,
		"exp":  time.Now().Add(duration).Unix(),
		"iat":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret()) // dynamically read
}

func VerifyJWT(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok && int64(exp) < time.Now().Unix() {
			return nil, jwt.ErrTokenExpired
		}
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
