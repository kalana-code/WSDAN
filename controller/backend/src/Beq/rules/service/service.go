package service

import (
	dispurserQueue "Beq/dispurser/db"
	JobModel "Beq/dispurser/model"
	"Beq/rules/db"
	"Beq/rules/model"
	sdb "Beq/settings/db"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var dispurserDb = dispurserQueue.GetRequestQueue()

// AddRule rules
func AddRule(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	log.Println("INFO: [RU]: Add Rule Process Initiate")
	rulesDb := db.GetRuleStore()
	var RuleData model.RulesDataRow
	reqBody, err := ioutil.ReadAll(r.Body)
	resp := model.Response{}
	// set Default value
	resp.Default()

	if err != nil {
		log.Println("ERROR: Payload Error", err)
		resp.BadRequest()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		err := json.Unmarshal(reqBody, &RuleData)
		if err != nil {
			log.Println("ERROR: Payload Error", err)
			resp.BadRequest()
			w.WriteHeader(http.StatusBadRequest)
		} else {
			log.Println("INFO: [RU]: Added Rule Successfully")
			rulesDb.AddRule(RuleData)
			var setting = sdb.GetSystemSetting()
			state, err := setting.IsForceDisposed()
			if state && err != nil {
				dispurserDb.AddJob(
					JobModel.Job{
						NodeIP:      RuleData.NodeIP,
						Type:        JobModel.TypeAddRule,
						TaskDetails: RuleData,
					},
				)
			}

			resp.Code = http.StatusOK
			resp.Message = "Data Base Updated"
			resp.Data = nil
			w.WriteHeader(http.StatusOK)
		}
	}
	json.NewEncoder(w).Encode(resp)

}

// GetAllRules for get all rules in controller
func GetAllRules(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	log.Println("INFO: [RU]: Get All Rule Process Initiate")
	rules := db.GetRuleStore()
	resp := model.Response{}
	// set Default value
	resp.Default()

	allRules, err := rules.GetAllRules()

	if err != nil {
		log.Println("ERROR: Payload Error", err)
		resp.InternalServerError()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		data := make(map[string]interface{})
		data["Rules"] = allRules
		resp.Code = http.StatusOK
		resp.Message = "Updated time: "
		resp.Data = data
		w.WriteHeader(http.StatusOK)
		log.Println("INFO: [RU]: Successfully Retrived All Rule Data")
	}
	json.NewEncoder(w).Encode(resp)

}

// RemoveRuleByRuleID remove rule ID
func RemoveRuleByRuleID(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	log.Println("INFO: [RU]: Remove Rule By RuleID Process Initiate")
	resp := model.Response{}
	// set Default value
	resp.Default()
	rules := db.GetRuleStore()
	RuleID := mux.Vars(r)["RuleID"]
	Message, NodeIP, err := rules.RemoveRuleByRuleID(RuleID)

	if err != nil {
		log.Println("ERROR: Payload Error", err)
		resp.InternalServerError()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		log.Println("INFO: [RU]: " + Message)
		if NodeIP != nil {
			jobModel := JobModel.Job{
				NodeIP: *NodeIP,
				Type:   JobModel.TypeRemoveRule,
				TaskDetails: JobModel.RemoveRuleJob{
					RuleID: RuleID,
				},
			}
			dispurserDb.AddJob(jobModel)
			err := rules.DispursedRule(RuleID)
			if err != nil {
				log.Println("ERROR: [RU]: ", err)
			}
		}
		resp.Code = http.StatusOK
		resp.Message = Message
		w.WriteHeader(http.StatusOK)

	}
	json.NewEncoder(w).Encode(resp)

}

// RemoveRulesByFlowID remove rules with given FlowID
func RemoveRulesByFlowID(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	log.Println("INFO: [RU]: Remove Rules By FlowID Process Initiate")
	resp := model.Response{}
	// set Default value
	resp.Default()
	rules := db.GetRuleStore()
	FlowID := mux.Vars(r)["FlowID"]
	Message, err := rules.RemoveRulesByFlowID(FlowID)

	if err != nil {
		log.Println("ERROR: Payload Error", err)
		resp.InternalServerError()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		log.Println("INFO: [RU]: Success Remove Rules By FlowID Process")
		resp.Code = http.StatusOK
		resp.Message = Message
		w.WriteHeader(http.StatusOK)

	}
	json.NewEncoder(w).Encode(resp)

}
