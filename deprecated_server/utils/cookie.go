package utils

import (
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func LoginFromCookie(r *http.Request) (string, error) {
	c, err := r.Cookie("token")
	if err != nil {
		return "", err
	}

	claims := &Claims{}
	tokenString := c.Value

	jwt.ParseWithClaims(tokenString, claims, func(tk *jwt.Token) (interface{}, error) {
		return os.Getenv("token_password"), nil
	})

	return claims.Login, nil
}

func SetTokenCookie(w http.ResponseWriter, login string) {
	expTime := ExpirationTime()
	claims := NewClaims(login, expTime)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   claims.String(),
		Expires: expTime,
		Path:    "/",

		HttpOnly: true,
	})
}

func DiscardTokenCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",

		HttpOnly: true,
	})
}
