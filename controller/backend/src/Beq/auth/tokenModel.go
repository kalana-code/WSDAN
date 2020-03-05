package auth

import "github.com/dgrijalva/jwt-go"

// Token is used handle jwt token
type Token struct {
	FirstName string
	LastName  string
	Gender    string
	Email     string
	Role      string
	*jwt.StandardClaims
}
