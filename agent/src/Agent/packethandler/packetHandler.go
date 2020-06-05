package packethandler

import (
	"Agent/database"
	"Agent/flowmanager"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var (
	ipl         *layers.IPv4   = &layers.IPv4{}
	icmpl       *layers.ICMPv4 = &layers.ICMPv4{}
	tcpl        *layers.TCP    = &layers.TCP{}
	udpl        *layers.UDP    = &layers.UDP{}
	dataPayload []byte         = nil
	buffer      gopacket.SerializeBuffer
	err         error
)

// PacketAnalyzer is used to anaylize the packet and create a new paccket
func PacketAnalyzer(db map[string]database.RuleConfiguration, packet gopacket.Packet) gopacket.SerializeBuffer {
	log.Println("Starting PacketAnalyzer ...!")
	packetDetails := flowmanager.PacketDetails{}
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		log.Println("IPv4 layer detected.")
		ipl, _ = ipLayer.(*layers.IPv4)
		packetDetails.DstIP = ipl.DstIP.String()
		packetDetails.Protocol = ipl.Protocol.String()
		fmt.Println(packetDetails.Protocol)
	}
	udpLayer := packet.Layer(layers.LayerTypeUDP)
	if udpLayer != nil {
		log.Println("UDP layer detected.")
		udpl, _ = udpLayer.(*layers.UDP)
	}
	icmpLayer := packet.Layer(layers.LayerTypeICMPv4)
	if icmpLayer != nil {
		log.Println("ICMP layer detected.")
		icmpl, _ = icmpLayer.(*layers.ICMPv4)
	}
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		log.Println("TCP layer detected.")
		tcpl, _ = tcpLayer.(*layers.TCP)
	}
	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		log.Println("Application layer detected.")
		dataPayload = applicationLayer.Payload()
	}
	ethernetLayer := generateEthernetLayer(db, packetDetails)
	options := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths:       true,
	}
	buffer = gopacket.NewSerializeBuffer()
	if udpLayer != nil {
		log.Println("UDP packet processing.")
		udpl.SetNetworkLayerForChecksum(ipl)
		err = gopacket.SerializeLayers(buffer, options,
			ethernetLayer,
			ipl,
			udpl,
			gopacket.Payload(dataPayload),
		)
		if err != nil {
			log.Fatal(err)
		}
	} else if icmpLayer != nil {
		log.Println("ICMP packet processing.")
		err = gopacket.SerializeLayers(buffer, options,
			ethernetLayer,
			ipl,
			icmpl,
			gopacket.Payload(dataPayload),
		)
		if err != nil {
			log.Fatal(err)
		}
	} else if tcpLayer != nil {
		log.Println("TCP packet processing.")
		tcpl.SetNetworkLayerForChecksum(ipl)
		err = gopacket.SerializeLayers(buffer, options,
			ethernetLayer,
			ipl,
			tcpl,
			gopacket.Payload(dataPayload),
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	return buffer
}

func generateEthernetLayer(db map[string]database.RuleConfiguration, packetDetails flowmanager.PacketDetails) *layers.Ethernet {
	log.Println("Generating Ethernet Layer ...!")
	ruleConfiguration := flowmanager.RuleChecker(db, packetDetails)
	_, hardwareAddrs, _ := getIPAndMAC()
	srcMAC, _ := net.ParseMAC(hardwareAddrs)
	dstMAC, _ := net.ParseMAC(ruleConfiguration.DstMAC)
	ethernetLayer := &layers.Ethernet{
		SrcMAC:       srcMAC,
		DstMAC:       dstMAC,
		EthernetType: layers.EthernetType(0x0800),
	}
	return ethernetLayer
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
