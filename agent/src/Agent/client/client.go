package client

import (
	"Agent/initializer"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
)

var (
	controllerIP string = "192.168.0.4"
	port         string = "8081"
	endPoint     string = "AddNodeInfo"
	infoLog      string = "INFO: [CL]:"
	errorLog     string = "ERROR: [CL]:"
)

// SendNodeData is used to send node data to the controller
func SendNodeData(nodeName string, nodeGroup string) {
	log.Println(infoLog, "Sending Node data")
	ip, mac, err := initializer.GetIPAndMAC()
	log.Println(infoLog, "IP :", ip, "MAC :", mac)
	if err != nil {
		log.Println(errorLog, "Not able to get IP and MAC :", err)
	} else {
		neighbours, _ := getNeighbours()
		jsonNodeData := map[string]interface{}{"NAME": nodeName, "GROUP": nodeGroup, "IP": ip, "MAC": mac}
		jsonNeighboursValues := make([]interface{}, len(neighbours))
		for i, nbMAC := range neighbours {
			linkBW, _ := getLinkThroughput(nbMAC)
			jsonNeighbourData := map[string]interface{}{"MAC": nbMAC, "Bandwidth": linkBW}
			jsonNeighboursValues[i] = jsonNeighbourData
		}
		jsonData := map[string]interface{}{"Node": jsonNodeData, "Neighbours": jsonNeighboursValues}
		jsonValue, _ := json.Marshal(jsonData)
		log.Println(infoLog, "Node Data:", string(jsonValue))
		dataSendingURL := "http://" + controllerIP + ":" + port + "/" + endPoint
		response, err := http.Post(dataSendingURL, "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			log.Println(errorLog, "The HTTP request failed with error", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			log.Println(infoLog, "Response Data:", string(data))
		}
	}
}

func getNeighbours() ([]string, error) {
	log.Println(infoLog, "Getting neighbour MAC addresses")
	out, err := exec.Command("sudo", "batctl", "n").Output()
	if err != nil {
		return nil, err
	}
	output := string(out[:])
	re := regexp.MustCompile(`([a-z0-9]+)\:([a-z0-9]+)\:([a-z0-9]+)\:([a-z0-9]+)\:([a-z0-9]+)\:([a-z0-9]+)`)
	match := re.FindAllString(output, -1)
	return match[2:], nil
}

func getLinkThroughput(nbMAC string) (string, error) {
	log.Println(infoLog, "Getting link throughput")
	// out, err := exec.Command("sudo", "batctl", "tp", nbMAC).Output()
	// if err != nil {
	// 	return "nil", err
	// }
	// output := string(out[:])
	// var rgx = regexp.MustCompile(`\((.*?)\)`)
	// rs := rgx.FindStringSubmatch(output)
	// return rs[1], nil
	throughput := rand.Intn(15) + 20
	throughputVal := strconv.Itoa(throughput) + " Mbps"
	return throughputVal, nil
}
