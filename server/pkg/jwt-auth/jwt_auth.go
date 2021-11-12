package jwt_auth

import (
	"github.com/go-chi/jwtauth"
)

func NewJWTAuth(secret string) *jwtauth.JWTAuth {
	return jwtauth.New("HS256", []byte(secret), nil)
}
