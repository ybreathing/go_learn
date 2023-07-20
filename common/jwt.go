package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"zzy/go-learn/module"
)

var jwtKey = []byte("zzy_secret_key")

type Claims struct {
	userId uint
	jwt.StandardClaims
}

func ReleaseToken(user module.User) (string, error) {
	expiresAt := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		userId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			Issuer:    "zzy",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
