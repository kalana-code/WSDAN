package server

import (
	"Agent/flowmanager"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	infoLog  string = "INFO: [SH]:"
	errorLog string = "ERROR: [SH]:"
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
	router.HandleFunc("/AddNodeInfo", addNodeRule)
	log.Println(http.ListenAndServe(":8081", router))
}

func addNodeRule(w http.ResponseWriter, r *http.Request) {
	log.Print(infoLog, "Adding Node Rule")
	rule := flowmanager.ControllerRuleConfiguration{}
	err := json.NewDecoder(r.Body).Decode(&rule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonData, jsonErr := json.Marshal(rule)
	if jsonErr != nil {
		log.Println(errorLog, "Recieved Rule Error:", jsonErr)
	}
	log.Print(infoLog, "Recieved Rule:", jsonData)
	flowmanager.RuleUpdater(rule)
}
