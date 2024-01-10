package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = "your_secret_key"

func GenerateJWT(isuser string) (string, error) {
	claims := jwt.MapClaims{
		"isuser": isuser,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SecretKey))
}

func ParseJWT(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	claims := token.Claims.(*jwt.MapClaims)
	return claims.GetIssuer()
}
