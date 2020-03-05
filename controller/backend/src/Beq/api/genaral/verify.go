package genaral

import (
	"Beq/auth"
	"encoding/json"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// Verify used when User login
func Verify(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	var header = r.Header.Get("x-access-token") //Grab the token from the header
	token := &auth.Token{}

	_, err := jwt.ParseWithClaims(header, token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("jwtSecret")), nil
	})
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}
	json.NewEncoder(w).Encode(token)

}
