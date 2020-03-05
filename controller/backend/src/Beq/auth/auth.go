package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// JwtVerify Middleware function
func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if (*r).Method == "OPTIONS" {
			return
		}

		var header = r.Header.Get("x-access-token") //Grab the token from the header

		header2 := strings.TrimSpace(header)
		log.Println(header2)
		if header2 == "" {
			// Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Unauthorized")
			return
		}
		tk := &Token{}

		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("jwtSecret")), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Unauthorized")
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
