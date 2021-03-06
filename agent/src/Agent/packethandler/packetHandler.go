package packethandler

import (
	"Agent/flowmanager"
	"Agent/initializer"
	"log"
	"net"

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
	infoLog     string = "INFO: [PH]:"
	errorLog    string = "ERROR: [PH]:"
)

// PacketAnalyzer is used to anaylize an input packet and create a new output packet
func PacketAnalyzer(packet gopacket.Packet) gopacket.SerializeBuffer {
	log.Println(infoLog, "Starting PacketAnalyzer")
	packetDetails := flowmanager.PacketDetails{}
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		log.Println(infoLog, "IPv4 layer detected.")
		ipl, _ = ipLayer.(*layers.IPv4)
		packetDetails.DstIP = ipl.DstIP.String()
		packetDetails.SrcIP = ipl.SrcIP.String()
		packetDetails.Protocol = ipl.Protocol.String()
	}
	udpLayer := packet.Layer(layers.LayerTypeUDP)
	if udpLayer != nil {
		log.Println(infoLog, "UDP layer detected.")
		udpl, _ = udpLayer.(*layers.UDP)
	}
	icmpLayer := packet.Layer(layers.LayerTypeICMPv4)
	if icmpLayer != nil {
		log.Println(infoLog, "ICMP layer detected.")
		icmpl, _ = icmpLayer.(*layers.ICMPv4)
	}
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		log.Println(infoLog, "TCP layer detected.")
		tcpl, _ = tcpLayer.(*layers.TCP)
	}
	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		log.Println(infoLog, "Application layer detected.")
		dataPayload = applicationLayer.Payload()
	}
	ethernetLayer := generateEthernetLayer(packetDetails)
	if ethernetLayer != nil {
		options := gopacket.SerializeOptions{
			ComputeChecksums: true,
			FixLengths:       true,
		}
		buffer = gopacket.NewSerializeBuffer()
		if udpLayer != nil {
			log.Println(infoLog, "UDP packet processing.")
			udpl.SetNetworkLayerForChecksum(ipl)
			err = gopacket.SerializeLayers(buffer, options,
				ethernetLayer,
				ipl,
				udpl,
				gopacket.Payload(dataPayload),
			)
			if err != nil {
				log.Println(errorLog, "UDP packet serilizing", err)
			}
		} else if icmpLayer != nil {
			log.Println(infoLog, "ICMP packet processing.")
			err = gopacket.SerializeLayers(buffer, options,
				ethernetLayer,
				ipl,
				icmpl,
				gopacket.Payload(dataPayload),
			)
			if err != nil {
				log.Println(errorLog, "ICMP packet serilizing", err)
			}
		} else if tcpLayer != nil {
			log.Println(infoLog, "TCP packet processing.")
			tcpl.SetNetworkLayerForChecksum(ipl)
			err = gopacket.SerializeLayers(buffer, options,
				ethernetLayer,
				ipl,
				tcpl,
				gopacket.Payload(dataPayload),
			)
			if err != nil {
				log.Println(errorLog, "TCP packet serilizing", err)
			}
		}
		return buffer
	}
	return nil
}

func generateEthernetLayer(packetDetails flowmanager.PacketDetails) *layers.Ethernet {
	log.Println(infoLog, "Generating Ethernet Layer")
	ruleConfiguration := flowmanager.RuleChecker(packetDetails)
	if ruleConfiguration.Action == "ACCEPT" {
		_, hardwareAddrs, _ := initializer.GetIPAndMAC()
		srcMAC, _ := net.ParseMAC(hardwareAddrs)
		dstMAC, _ := net.ParseMAC(ruleConfiguration.DstMAC)
		ethernetLayer := &layers.Ethernet{
			SrcMAC:       srcMAC,
			DstMAC:       dstMAC,
			EthernetType: layers.EthernetType(0x0800),
		}
		return ethernetLayer
	}
	return nil
}
