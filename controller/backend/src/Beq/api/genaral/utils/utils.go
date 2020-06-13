package utils

import (
	"Beq/api/genaral/model"
	"Beq/auth"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// import dataBase
var db = GetUserStore()

// InsertUser function save user in db
func InsertUser(User model.UserInfo) model.InnerResponse {

	errIns := db.AddUser(User)

	if errIns != nil {
		log.Println("ERROR: DataBase Insert Query Excuting Process Failed. ", errIns)
		if strings.Contains(errIns.Error(), "Email_UNIQUE") {
			return errorDuplicateEmail
		}
		return errorDataBase
	}

	log.Println("INFO: User has been inserted successfully.")
	return stateCreated

}

// FindUser used for find a user
func FindUser(User model.UserLogin) model.InnerResponse {
	userInfo := &model.UserInfo{}
	resp := model.InnerResponse{}
	userInfo, err := db.FindUser(User.Email)
	if err != nil {
		log.Println("ERROR: Email address not found", err)
		resp.Status = http.StatusForbidden
		resp.Message = "Email address not found"
		return resp
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errEncry := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(User.Password))
	if errEncry != nil && errEncry == bcrypt.ErrMismatchedHashAndPassword {
		log.Println("ERROR: Invalid login credentials. Please try again", errEncry)
		resp.Status = http.StatusForbidden
		resp.Message = "Invalid login credentials. Please try again"
		return resp
	}

	JwtToken := &auth.Token{
		FirstName: userInfo.FirstName,
		LastName:  userInfo.LastName,
		Gender:    userInfo.Gender,
		Email:     userInfo.Email,
		Role:      userInfo.Role,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), JwtToken)
	// sign token
	tokenString, error := token.SignedString([]byte(os.Getenv("jwtSecret")))

	if error != nil {
		log.Println("ERROR: Token signing process failed", error)
		resp.Status = http.StatusInternalServerError
		resp.Message = "token signing process failed"
		return resp
	}
	resp.Status = http.StatusOK
	resp.Message = "logged in successfully"
	resp.Data = map[string]interface{}{"token": tokenString}
	log.Println("INFO: Logged in successfully")

	return resp
}
