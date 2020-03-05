package genaral

import (
	"Beq/api/genaral/model"
	"Beq/api/genaral/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Register used for register a new user
func Register(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	var TempUser model.User
	reqBody, err := ioutil.ReadAll(r.Body)
	// message on Response
	resp := model.Response{}
	// set Default value
	resp.Default()

	if err != nil {
		log.Println("ERROR: Payload Error . ", err)
		resp.BadRequest()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		err2 := json.Unmarshal(reqBody, &TempUser)
		if err2 == nil {
			pass, errCrpto := bcrypt.GenerateFromPassword([]byte(TempUser.Password), bcrypt.DefaultCost)
			if errCrpto != nil {
				log.Println("ERROR: Decryption process has been failed. ", errCrpto)
				resp.InternalServerError()
				w.WriteHeader(http.StatusInternalServerError)

			} else {
				// set encripted passward
				TempUser.Password = string(pass)
				StateMessage := utils.InsertUser(TempUser)
				// set response message
				resp.Code = StateMessage.Status
				resp.Message = StateMessage.Message

				w.WriteHeader(StateMessage.Status)
				json.NewEncoder(w).Encode(resp)
			}
		} else {
			log.Println("ERROR: Payload Error . ", err)
			resp.BadRequest()
			w.WriteHeader(http.StatusBadRequest)
		}

	}

}
