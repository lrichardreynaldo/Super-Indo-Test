package utils

import (
	"errors"
	models "superIndo/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("secret")

func ValidateToken(signedToken string) (userId int, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&models.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*models.Claims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return int(claims.UserId), nil
}
