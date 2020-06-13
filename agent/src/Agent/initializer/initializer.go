package initializer

import (
	"log"
	"net"
	"os/exec"
	"strings"
)

var (
	infoLog      string = "INFO: [IZ]:"
	errorLog     string = "ERROR: [IZ]:"
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

// IptableInitializer is used to intialize the iptables of the node
func IptableInitializer() error {
	log.Println(infoLog, "Run IptableInitializer")
	err := iptableRuleFlusher()
	if err != nil {
		log.Println(errorLog, "Error when flushing iptables")
		return err
	}
	err = iptableInputHandler()
	if err != nil {
		log.Println(errorLog, "Error when handling INPUT in iptables")
		return err
	}
	err = iptableOutputHandler()
	if err != nil {
		log.Println(errorLog, "Error when handling OUTPUT in iptables")
	}
	return err
}

// iptableRuleFlusher is used to flush iptable rules
func iptableRuleFlusher() error {
	log.Println(infoLog, "Flushing iptable rules")
	return exec.Command("sudo", "iptables", "-F").Run()
}

// iptableInputHandler is used to insert iptable INPUT rules
func iptableInputHandler() error {
	log.Println(infoLog, "Inserting iptable INPUT rules")
	ip, _, _ := GetIPAndMAC()
	err := exec.Command("sudo", "iptables", "-A", "INPUT", "1", "-p", "tcp", "--dport", "22", "-j", "ACCEPT").Run()
	if err != nil {
		log.Println(errorLog, "Error when inserting INPUT rule: ssh")
		return err
	}
	err = exec.Command("sudo", "iptables", "-A", "INPUT", "2", "-d", ip, "-j", "ACCEPT").Run()
	if err != nil {
		log.Println(errorLog, "Error when inserting INPUT rule: allow packet with dst ip as node ip")
		return err
	}
	return exec.Command("sudo", "iptables", "-A", "INPUT", "3", "-i", "wlan0", "-j", "NFQUEUE", "--queue-num", "0").Run()
}

// iptableOutputHandler is used to insert iptable OUTPUT rules
func iptableOutputHandler() error {
	log.Println(infoLog, "Inserting iptable OUTPUT rules")
	err := exec.Command("sudo", "iptables", "-A", "OUTPUT", "-p", "tcp", "--sport", "22", "-j", "ACCEPT").Run()
	if err != nil {
		log.Println(errorLog, "Error when inserting OUTPUT rule: ssh")
		return err
	}
	err = exec.Command("sudo", "iptables", "-A", "OUTPUT", "-d", controllerIP, "-j", "ACCEPT").Run()
	if err != nil {
		log.Println(errorLog, "Error when inserting OUTPUT rule: allow packet with dst ip as controller ip")
		return err
	}
	return exec.Command("sudo", "iptables", "-A", "OUTPUT", "-j", "wlan0", "-j", "NFQUEUE", "--queue-num", "0").Run()
}
