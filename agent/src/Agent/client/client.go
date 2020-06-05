package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)

var (
	controllerIP string = "192.168.0.4"
	port         string = "8081"
	endPoint     string = "AddNodeInfo"
)

// SendNodeData is used to send node data to the controller
func SendNodeData() {
	ip, mac, err := getIPAndMAC()

	if err != nil {
		fmt.Println("Not able to get IP and MAC : ", err)
	} else {
		neighbours, _ := getNeighbours()
		jsonNodeData := map[string]interface{}{"NAME": "Node", "GROUP": "Gateway", "IP": ip, "MAC": mac}
		jsonNeighboursValues := make([]interface{}, len(neighbours))
		for i, nbMAC := range neighbours {
			linkBW, _ := getLinkThroughput(nbMAC)
			jsonNeighbourData := map[string]interface{}{"MAC": nbMAC, "Bandwidth": linkBW}
			jsonNeighboursValues[i] = jsonNeighbourData
		}
		jsonData := map[string]interface{}{"Node": jsonNodeData, "Neighbours": jsonNeighboursValues}
		jsonValue, _ := json.Marshal(jsonData)
		fmt.Println(string(jsonValue))
		dataSendingURL := "http://" + controllerIP + ":" + port + "/" + endPoint
		response, err := http.Post(dataSendingURL, "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
	}
}

func getIPAndMAC() (string, string, error) {
	var currentIP, currentNetworkHardwareName string
	currentNetworkHardwareName = "wlan0"
	netInterface, err := net.InterfaceByName(currentNetworkHardwareName)

	if err != nil {
		return "nil", "nil", err
	}

	macAddress := netInterface.HardwareAddr
	addresses, err := netInterface.Addrs()
	currentIP = addresses[0].String()
	ipAddr := currentIP[:strings.IndexByte(currentIP, '/')]
	hwAddr, err := net.ParseMAC(macAddress.String())

	if err != nil {
		fmt.Println("No able to parse MAC address : ", err)
		return "nil", "nil", err
	}

	return ipAddr, hwAddr.String(), nil
}

func getNeighbours() ([]string, error) {
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
	out, err := exec.Command("sudo", "batctl", "tp", nbMAC).Output()

	if err != nil {
		return "nil", err
	}

	output := string(out[:])
	var rgx = regexp.MustCompile(`\((.*?)\)`)
	rs := rgx.FindStringSubmatch(output)
	return rs[1], nil
}
