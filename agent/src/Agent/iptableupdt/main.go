package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type NodeDetails struct {
	GROUP string `json:"GROUP"`
	IP    string `json:"IP"`
	MAC   string `json:"MAC"`
	NAME  string `json:"NAME"`
}

var a = 0
var m = make(map[int]string)

func main() {
	fmt.Print("Start initialization !!!\n")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/sendIptableInfo", homeLink)
	log.Fatal(http.ListenAndServe(":8081", router))
	fmt.Println("Starting the application...")
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved !!!")
	nodeDet := NodeDetails{}
	err := json.NewDecoder(r.Body).Decode(&nodeDet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonData, jsonErr := json.Marshal(nodeDet)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	//fmt.Println(string(jsonData))
	m[a] = string(jsonData)

	if a == 2 {
		for key, val := range m {
			fmt.Println("key: ", key, ", val: ", val)
		}
	}
	a++

}
