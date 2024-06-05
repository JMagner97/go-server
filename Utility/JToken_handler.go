package utility

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtToken = []byte("1OE2Typ4QmreYCYvX1V1uJEZHZVy6mSo")

type UserAuthentication struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWTToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Hour)
	userAuth := &UserAuthentication{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "your app name",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userAuth)
	tokenString, err := token.SignedString(jwtToken)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func IsValidToken(tokenString string) (*UserAuthentication, error, bool) {
	userAuth := &UserAuthentication{}
	token, err := jwt.ParseWithClaims(tokenString, userAuth, func(token *jwt.Token) (interface{}, error) { return jwtToken, nil })

	if err != nil {
		return nil, err, false
	}
	if !token.Valid {
		return nil, fmt.Errorf("Invalid token"), false
	}
	return userAuth, nil, true
}
