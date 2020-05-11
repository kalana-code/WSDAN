package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

type RuleConfiguration struct {
	SOURCE string `json:"Source"`
	ACTION string `json:"Action"`
}

var a = 0
var m = make(map[int]RuleConfiguration)
var rules []RuleConfiguration

func main() {
	fmt.Print("Start initialization !!!\n")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/sendIptableInfo", homeLink)
	log.Fatal(http.ListenAndServe(":8081", router))
	fmt.Println("Starting the application...")
}

func homeLink(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Recieved !!!")
	ruleConf := RuleConfiguration{}
	err := json.NewDecoder(r.Body).Decode(&ruleConf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// jsonData, jsonErr := json.Marshal(ruleConf)
	// if jsonErr != nil {
	// 	log.Fatal(jsonErr)
	// }

	//fmt.Println(string(jsonData))
	m[a] = ruleConf

	// if a == 2 {
	for key, val := range m {
		fmt.Println("key: ", key, ", val: ", val)

		cmd1 := exec.Command("sudo", "iptables", "--flush")
		_, err1 := cmd1.CombinedOutput()
		if err1 != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err1)
		}

		cmd2 := exec.Command("sudo", "iptables", "-A", "INPUT", "-i", "eth0", "-p", "tcp", "--dport", "22", "-j", "ACCEPT")
		_, err2 := cmd2.CombinedOutput()
		if err2 != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err2)
		}

		cmd3 := exec.Command("sudo", "iptables", "-A", "INPUT", "-i", "wlan0", "-s", val.SOURCE, "-j", val.ACTION)
		_, err3 := cmd3.CombinedOutput()
		if err3 != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err3)
		}

		cmd4 := exec.Command("sudo", "iptables", "-A", "INPUT", "-j", "DROP")
		_, err4 := cmd4.CombinedOutput()
		if err4 != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err4)
		}
	}

	// for _, rule := range rules {
	// 	action := rule.ACTION
	// 	// do something with it? I don't know
	// 	fmt.Println(action) // I guess?
	// }
	// }

	//fmt.Printf("combined out:\n%s\n", string(out))

	a++

}
