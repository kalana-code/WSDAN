package main

<<<<<<< HEAD
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
	"time"
)

func sendNodeData() {
	ip, mac, err := getIPAndMAC()

	if err != nil {
		fmt.Println("Not able to get IP and MAC : ", err)
	} else {
		//neighbours, _ := getNeighbours()
		jsonNodeData := map[string]interface{}{"NAME": "Node", "GROUP": "Gateway", "IP": ip, "MAC": mac}
		jsonNeighboursValues := make([]interface{}, 1)
		// for i, nbMAC := range neighbours {
		jsonNeighbourData := map[string]interface{}{"MAC": "93:FB:E5:3D:0E:CF", "Bandwidth": 10}
		jsonNeighboursValues[0] = jsonNeighbourData
		// }
		jsonData := map[string]interface{}{"Node": jsonNodeData, "Neighbours": jsonNeighboursValues}
		jsonValue, _ := json.Marshal(jsonData)
		fmt.Println(string(jsonValue))

		response, err := http.Post("http://192.168.8.100:8081/AddNodeInfo", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
	}
}

func doEvery(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}

func main() {
	doEvery(5000*time.Millisecond, sendNodeData)
}

func getIPAndMAC() (string, string, error) {
	var currentIP, currentNetworkHardwareName string
	currentNetworkHardwareName = "en0"
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
=======
import "fmt"

func main() {
	fmt.Print("Kalana")
>>>>>>> f298c4eec634aa21bde2a5be0d4bf91e78d5cc14
}
