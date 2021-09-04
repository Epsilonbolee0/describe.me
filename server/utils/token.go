package utils

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}

func ExpirationTime() time.Time {
	return time.Now().Add(30 * time.Minute)
}

func NewClaims(login string, expTime time.Time) *Claims {
	return &Claims{
		Login: login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
}

func (claims *Claims) HasExpired() bool {
	return time.Now().Unix()-claims.ExpiresAt > 30
}

func (claims *Claims) String() string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("secret_password")))
	if err != nil {
		panic(err)
	}

	return tokenString
}
