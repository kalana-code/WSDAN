package genaral

import (
	"Beq/api/genaral/model"
	"Beq/api/genaral/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Login used when User login
func Login(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	var user model.UserLogin
	reqBody, err := ioutil.ReadAll(r.Body)
	resp := model.Response{}
	// set Default value
	resp.Default()

	if err != nil {
		log.Println("ERROR: Payload Error", err)
		resp.BadRequest()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		err := json.Unmarshal(reqBody, &user)
		if err != nil {
			log.Println("ERROR: Payload Error", err)
			resp.BadRequest()
			w.WriteHeader(http.StatusBadRequest)
		} else {
			StateMessage := utils.FindUser(user)
			resp.Code = StateMessage.Status
			resp.Message = StateMessage.Message
			resp.Data = StateMessage.Data
			w.WriteHeader(StateMessage.Status)
		}
	}
	json.NewEncoder(w).Encode(resp)

}
