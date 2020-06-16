package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	device      string = "wlan0"
	snapshotLen int32  = 1024
	promiscuous bool   = false
	err         error
	timeout     time.Duration = 1 * time.Second
	handle      *pcap.Handle
	buffer      gopacket.SerializeBuffer
	options     gopacket.SerializeOptions = gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths:       true,
	}
	ethernetLayer *layers.Ethernet = &layers.Ethernet{}
	ethl          *layers.Ethernet = &layers.Ethernet{}
	ipl           *layers.IPv4     = &layers.IPv4{}
	tcpl          *layers.TCP      = &layers.TCP{}
	payload       []byte           = nil
)

func main() {
	// Open device
	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set filter
	var filter string = "icmp"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Only capturing TCP port 80 packets.")

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {

		// Create a properly formed packet, just with
		// empty details. Should fill out MAC addresses,
		// IP addresses, etc.
		// This time lets fill out some information
		// ipLayer := &layers.IPv4{
		// 	SrcIP: net.IP{127, 0, 0, 1},
		// 	DstIP: net.IP{8, 8, 8, 8},
		// }
		ethernetLayerin := packet.Layer(layers.LayerTypeEthernet)
		// ethernetLayer.SrcMAC = net.HardwareAddr{0xB8, 0x27, 0xEB, 0x81, 0x12, 0xEB}
		// ethernetLayer.DstMAC = net.HardwareAddr{0xB8, 0x27, 0xEB, 0x9A, 0x5E, 0xA5}
		if ethernetLayerin != nil {
			ethl, _ = ethernetLayerin.(*layers.Ethernet)
			ethernetLayer = &layers.Ethernet{
				SrcMAC:       net.HardwareAddr{0xB8, 0x27, 0xEB, 0x81, 0x12, 0xEB},
				DstMAC:       net.HardwareAddr{0xB8, 0x27, 0xEB, 0x9A, 0x5E, 0xA5},
				EthernetType: ethl.EthernetType,
				//	BaseLayer:       ethl.BaseLayer,
				//	CanDecode:       ethl.CanDecode,
				//	DecodeFromBytes: ethl.DecodeFromBytes,
				//	LayerContents:   ethl.LayerContents,
				//	LayerPayload:    ethl.LayerPayload,
				//	LayerType:       ethl.LayerType,
				Length: ethl.Length,
				//	LinkFlow:        ethl.LinkFlow,
				//	NextLayerType:   ethl.NextLayerType,
				//	Payload:         ethl.Payload,
				//	SerializeTo:     ethl.SerializeTo,
			}
		}

		// tcpLayer := &layers.TCP{
		// 	SrcPort: layers.TCPPort(4321),
		// 	DstPort: layers.TCPPort(80),
		// }
		// And create the packet with the layers

		// Let's see if the packet is IP (even though the ether type told us)
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		if ipLayer != nil {
			fmt.Println("IPv4 layer detected.")
			ipl, _ = ipLayer.(*layers.IPv4)
		}

		// Let's see if the packet is TCP
		tcpLayer := packet.Layer(layers.Layer.LayerTypeTCP)
		if tcpLayer != nil {
			fmt.Println("TCP layer detected.")
			tcpl, _ = tcpLayer.(*layers.TCP)
			tcpl.SYN
		}

		udpLayer := packet.Layer(layers.Layer.LayerTypeUDP)
		if udpLayer != nil {
			fmt.Println("UDP layer detected.")
			udpl, _ = udpLayer.(*layers.UDP)

		}

		applicationLayer := packet.ApplicationLayer()
		if applicationLayer != nil {

			// Search for a string inside the payload
			if strings.Contains(string(applicationLayer.Payload()), "HTTP") {
				fmt.Println("HTTP found!")
			}
			payload = applicationLayer.Payload()
		}
		buffer = gopacket.NewSerializeBuffer()

		tcpl.SetNetworkLayerForChecksum(ipl)
		err = gopacket.SerializeLayers(buffer, options,
			ethernetLayer,
			ipl,
			tcpl,
			gopacket.Payload(payload),
		)
		if err != nil {
			panic(err)
		}
		outgoingPacket := buffer.Bytes()

		err = handle.WritePacketData(outgoingPacket)
		if err != nil {
			log.Fatal(err)
		}
	}
}
