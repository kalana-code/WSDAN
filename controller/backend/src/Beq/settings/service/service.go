package service

import (
	"Beq/nodes/model"
	"Beq/settings/db"
	"encoding/json"
	"log"
	"net/http"
	// "github.com/openshift/geard/containers/http"
)

//Toggle  toggle system state
func Toggle(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	log.Println("INFO: [ST]: Toggling System Mode Process is Initiated.")
	setting := db.GetSystemSetting()
	resp := model.Response{}
	// set Default value
	resp.Default()

	err := setting.ToggleMode()

	if err != nil {
		log.Println("ERROR: [ST]: Toggle Process is Failed.", err)
		resp.InternalServerError()
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		resp.Code = http.StatusOK
		resp.Message = "System State is Toggled Successfully."
		w.WriteHeader(http.StatusOK)
		log.Println("INFO: [ST]: System State is Toggled Successfully.")
	}
	json.NewEncoder(w).Encode(resp)

}

//ToggleForceDispurserMode  toggle system state
func ToggleForceDispurserMode(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	log.Println("INFO: [ST]: Toggling Force Dispurser Mode Process is Initiated.")
	setting := db.GetSystemSetting()
	resp := model.Response{}
	// set Default value
	resp.Default()

	err := setting.ToggleForceDispurserMode()

	if err != nil {
		log.Println("ERROR: [ST]: Toggle Process is Failed.", err)
		resp.InternalServerError()
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		resp.Code = http.StatusOK
		resp.Message = "Toggling Force Dispurser Mode is Toggled Successfully."
		w.WriteHeader(http.StatusOK)
		log.Println("INFO: [ST]: Toggling Force Dispurser Mode is Toggled Successfully.")
	}
	json.NewEncoder(w).Encode(resp)

}

//GetCurrentSetting  used for get current setting of system
func GetCurrentSetting(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	log.Println("INFO: [ST]: System Current Setting is  requested.")
	setting := db.GetSystemSetting()
	resp := model.Response{}
	// set Default value
	resp.Default()

	currentSetting, err := setting.GetSetting()

	if err != nil {
		log.Println("ERROR: [ST]: System current setting requesting process is  failed.", err)
		resp.InternalServerError()
		w.WriteHeader(http.StatusInternalServerError)
	} else {

		resp.Code = http.StatusOK
		resp.Data = *currentSetting
		resp.Message = "System current setting is  recived  successfully."
		w.WriteHeader(http.StatusOK)
		log.Println("INFO: [ST]: System current setting is  sent  successfully.")
	}
	json.NewEncoder(w).Encode(resp)

}
