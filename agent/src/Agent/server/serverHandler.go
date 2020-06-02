package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

// func main() {
// 	fmt.Print("Start initialization !!!\n")
// 	router := mux.NewRouter().StrictSlash(true)
// 	router.HandleFunc("/AddNodeInfo", homeLink)
// 	log.Fatal(http.ListenAndServe(":8081", router))
// 	fmt.Println("Starting the application...")
// }

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved !!!")
	nodeDet := NodeData{}
	err := json.NewDecoder(r.Body).Decode(&nodeDet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonData, jsonErr := json.Marshal(nodeDet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(string(jsonData))

}
