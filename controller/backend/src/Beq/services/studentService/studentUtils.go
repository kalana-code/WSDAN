package studentservice

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"Beq/auth"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func findStudent(email, password string, w http.ResponseWriter, db *sql.DB) map[string]interface{} {
	student := &studentLoginModal{}
	err := db.QueryRow("SELECT Email, Password FROM Students where Email = ?", email).Scan(&student.UserName, &student.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: data base :", err)
		var resp = map[string]interface{}{"Message": "Email address not found"}
		return resp
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		w.WriteHeader(http.StatusInternalServerError)
		var resp = map[string]interface{}{"Message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := &auth.Token{
		UserName: student.UserName,
		Role:     "STUDENT",
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte(os.Getenv("jwtSecret")))
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		var resp = map[string]interface{}{"Message": "Encription Process Failed"}
		return resp
	}

	var resp = map[string]interface{}{"Message": "logged in successfully"}
	resp["token"] = tokenString
	resp["Profile"] = "data:image/png;base64,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"

	//Store the token in the response
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(time.Minute * 5),
	})
	w.WriteHeader(http.StatusOK)
	return resp
}
