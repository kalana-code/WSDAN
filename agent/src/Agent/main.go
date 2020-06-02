package main

import (
	"Agent/database"
	"Agent/input"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	device        string = "wlan0"
	snapshotLen   int32  = 1024
	promiscuous   bool   = false
	err           error
	timeout       time.Duration = 30 * time.Second
	handle        *pcap.Handle
	ruleConf      *database.RuleConfiguration
	found         bool
	encapPacket   input.EncapsulatedPacket
	controllerIP  string = "192.168.0.4"
	controllerMAC string = "B8:27:EB:9A:5E:A5"
)

func main() {
	db := database.CreateDatabase()
	newRule1 := database.RuleConfiguration{EnacapPcktDstIP: "192.168.0.4", EnacapPcktDstMAC: "b8:27:eb:9a:5e:a5", Action: "ACCEPT"}
	database.CreateRule(db, "192.168.0.5", newRule1)
	database.ViewRules(db)

	ipAddr, macAddr, err := getIPAndMAC()
	if err != nil {
		log.Fatal(err)
	}
	// Open device
	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		pcktDesctinationIP := input.GetPacketDstIP(packet)
		ruleConf, found = database.CheckRule(db, pcktDesctinationIP)
		if found {
			encapPacket = input.EncapsulatedPacket{
				SrcIP:  ipAddr,
				DstIP:  ruleConf.EnacapPcktDstIP,
				SrcMAC: macAddr,
				DstMAC: ruleConf.EnacapPcktDstMAC,
				Packet: packet,
			}
		} else {
			encapPacket = input.EncapsulatedPacket{
				SrcIP:  ipAddr,
				DstIP:  controllerIP,
				SrcMAC: macAddr,
				DstMAC: controllerMAC,
				Packet: packet,
			}
		}
		input.CreateAndSendEncapsulatedPacket(handle, encapPacket)
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

// func main() {
// 	db := database.CreateDatabase()

// 	newRule1 := database.RuleConfiguration{EnacapPcktDstIP: "192.168.0.4", EnacapPcktDstMAC: "b8:27:eb:9a:5e:a5", Action: "ACCEPT"}
// 	database.CreateRule(db, "192.168.0.5", newRule1)
// 	database.ViewRules(db)
// 	api.SendNodeData("192.168.0.4", "8081", "AddNodeInfo")
// 	// func main() {
// 	// 	doEvery(5000*time.Millisecond, sendNodeData)
// 	// }

// }

// // doEvery is used to execute a function periodically
// func doEvery(d time.Duration, f func()) {
// 	for range time.Tick(d) {
// 		f()
// 	}
// }

// Set filter
// var filter string = "icmp"
// err = handle.SetBPFFilter(filter)
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Println("Only capturing ICMP packets")
// ip layer for forwarding packet
