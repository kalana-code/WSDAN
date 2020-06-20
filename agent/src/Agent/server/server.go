package server

import (
	"Agent/flowmanager"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	infoLog                string = "INFO: [SH]:"
	errorLog               string = "ERROR: [SH]:"
	port                   string = ":8082"
	endPointAddFlowRule    string = "/AddFlowRule"
	endPointRemoveFlowRule string = "/RemoveFlowRule"
)

//NodeDetails is used to store node data
type NodeDetails struct {
	GROUP string `json:"GROUP"`
	IP    string `json:"IP"`
	MAC   string `json:"MAC"`
	NAME  string `json:"NAME"`
}

//NodeNeighboursDetails is used to store neighbour data
type NodeNeighboursDetails struct {
	Bandwidth string `json:"Bandwidth"`
	MAC       string `json:"MAC"`
}

//NodeData is used to store both node and neighbour data
type NodeData struct {
	Node       NodeDetails             `json:"Node"`
	Neighbours []NodeNeighboursDetails `json:"Neighbours"`
}

// Server is used to retrieve data from the controller
func Server() {
	log.Print(infoLog, "Starting Server")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(endPointAddFlowRule, addNodeRule)
	router.HandleFunc(endPointRemoveFlowRule, removeNodeRule)
	log.Println(http.ListenAndServe(port, router))
}

func addNodeRule(w http.ResponseWriter, r *http.Request) {
	log.Print(infoLog, "Adding Node Rule")
	rule := flowmanager.ControllerRuleConfiguration{}
	err := json.NewDecoder(r.Body).Decode(&rule)
	if err != nil {
		log.Println(errorLog, "Recieved Rule Error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		log.Print(infoLog, "Rule is received successfully")
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(http.StatusText(http.StatusAccepted) + ": Rule is received successfully"))
		jsonData, jsonErr := json.Marshal(rule)
		if jsonErr != nil {
			log.Println(errorLog, "Recieved Rule JSON Error:", jsonErr)
		}
		log.Print(infoLog, "Recieved Rule:", string(jsonData))
		flowmanager.RuleUpdater(rule)
	}
}

func removeNodeRule(w http.ResponseWriter, r *http.Request) {
	log.Print(infoLog, "Removing Node Rule(By RuleID)")
	rule := flowmanager.RemoveRule{}
	err := json.NewDecoder(r.Body).Decode(&rule)
	if err != nil {
		log.Println(errorLog, "Recieved Remove Rule(By RuleID) Error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		log.Print(infoLog, "Remove rule(By RuleID) is received successfully")
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(http.StatusText(http.StatusAccepted) + ": Remoce Rule(By RuleID) is received successfully"))
		jsonData, jsonErr := json.Marshal(rule)
		if jsonErr != nil {
			log.Println(errorLog, "Recieved Remove Rule(By RuleID) JSON Error:", jsonErr)
		}
		log.Print(infoLog, "Recieved Remove Rule(By RuleID):", string(jsonData))
		flowmanager.RuleRemoveByRuleID(rule)
	}
}
