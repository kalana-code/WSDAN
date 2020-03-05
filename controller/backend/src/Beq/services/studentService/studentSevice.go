package studentservice

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"

// 	"golang.org/x/crypto/bcrypt"

// 	"Beq/utils"
// )

// var db = utils.ConnectDataBase()

// // RegisterStudent used for register a new student
// func RegisterStudent(w http.ResponseWriter, r *http.Request) {
// 	if (*r).Method == "OPTIONS" {
// 		return
// 	}
// 	var newStudent studentModal
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Println("ERROR: Payload Error . ", err)
// 		w.WriteHeader(http.StatusBadRequest)
// 	} else {
// 		err2 := json.Unmarshal(reqBody, &newStudent)
// 		if err2 == nil {
// 			pass, errCrpto := bcrypt.GenerateFromPassword([]byte(newStudent.Password), bcrypt.DefaultCost)
// 			if errCrpto != nil {
// 				log.Println("ERROR: Decryption process has been failed. ", errCrpto)
// 				w.WriteHeader(http.StatusInternalServerError)
// 			} else {

// 				newStudent.Password = string(pass)
// 				insForm, err := db.Prepare("INSERT INTO Students(FirstName,LastName, Email,Password,Gender,BirthDay) VALUES(?,?,?,?,?,?)")
// 				// executing
// 				if err != nil {
// 					log.Println("ERROR: DB insert query preparation process failed.", err)
// 					w.WriteHeader(http.StatusInternalServerError)
// 				} else {
// 					_, err := insForm.Exec(
// 						newStudent.FirstName,
// 						newStudent.LastName,
// 						newStudent.Email,
// 						newStudent.Password,
// 						newStudent.Gender,
// 						newStudent.BirthDay)
// 					if err != nil {
// 						log.Println("ERROR: DataBase Insert Query Excuting Process Failed. ")
// 						w.WriteHeader(http.StatusInternalServerError)
// 					} else {
// 						log.Println("INFO: User (Student) has been created.")
// 						w.WriteHeader(http.StatusCreated)
// 					}

// 				}
// 				defer insForm.Close()

// 			}
// 		} else {
// 			log.Println(err2)
// 			w.WriteHeader(http.StatusBadRequest)
// 		}

// 	}

// }

// // LoginStudent used for authorized login Used
// func LoginStudent(w http.ResponseWriter, r *http.Request) {
// 	if (*r).Method == "OPTIONS" {
// 		return
// 	}
// 	var studentLogin studentLoginModal
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 	} else {
// 		err2 := json.Unmarshal(reqBody, &studentLogin)
// 		if err2 != nil {
// 			log.Println(err2)
// 			w.WriteHeader(http.StatusBadRequest)
// 		} else {
// 			resp := findStudent(studentLogin.UserName, studentLogin.Password, w, db)
// 			json.NewEncoder(w).Encode(resp)
// 		}
// 	}

// }

// // TestAPI Test Api
// func TestAPI(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "kalana")
// }
