package utils

import (
	"log"
	"net"
	"strings"
)

var (
	infoLog      string = "INFO: [PH]:"
	errorLog     string = "ERROR: [PH]:"
	controllerIP string = "192.168.0.4"
)

// GetIPAndMAC is used to get IP and MAC of the device
func GetIPAndMAC() (string, string, error) {
	log.Println(infoLog, "Getting source MAC address and IP address")
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
		log.Println(errorLog, "Not able to parse MAC address :", err)
		return "nil", "nil", err
	}
	return ipAddr, hwAddr.String(), nil
}
