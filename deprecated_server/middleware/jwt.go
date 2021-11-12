package middleware

import (
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"

	"../utils"
)

var JwtAuthenithication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		noAuth := []string{"/auth/register", "/auth/login"}
		requestPath := r.URL.Path
		for _, value := range noAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenString := c.Value
		claims := &utils.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(tk *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("secret_password")), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if claims.HasExpired() {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		utils.SetTokenCookie(w, claims.Login)
		next.ServeHTTP(w, r)
	})
}
