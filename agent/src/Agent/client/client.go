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

// controllerMAC is used to store controller MAC
type controllerMAC struct {
	MAC string `json:"MAC"`
}

// responseGetControllerMAC is used to store controller MAC response
type responseGetControllerMAC struct {
	Program string        `json:"Program"`
	Version string        `json:"Version"`
	Status  string        `json:"Status"`
	Code    string        `json:"Code"`
	Message string        `json:"Message"`
	Data    controllerMAC `json:"Data"`
}

var (
	controllerIP             string = "192.168.0.4"
	port                     string = "8081"
	endPoint                 string = "AddNodeInfo"
	endPointGetControllerMAC string = "GetControllerMac"
	infoLog                  string = "INFO: [CL]:"
	errorLog                 string = "ERROR: [CL]:"
	data                     responseGetControllerMAC
)

// GetControllerMAC is used to get Controller MAC address
func GetControllerMAC() (string, error) {
	log.Println(infoLog, "Get Controller MAC")
	getMACUrl := "http://" + controllerIP + ":" + port + "/" + endPointGetControllerMAC
	resp, err := http.Get(getMACUrl)
	if err != nil {
		log.Println(errorLog, "The HTTP request(Get Controller MAC) failed with error", err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(errorLog, "GetControllerMAC: Error in reading response", err)
		return "", err
	}
	log.Println(infoLog, "GetControllerMAC: Response Data:", string(body))
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		log.Println(errorLog, "GetControllerMAC: JSON Error:", jsonErr)
		return "", jsonErr
	}
	return data.Data.MAC, nil
}

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
			log.Println(errorLog, "The HTTP request(Send Node Data) failed with error", err)
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
