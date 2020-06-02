package input

import (
	"fmt"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	err          error
	packetBuffer gopacket.SerializeBuffer
	encapBuffer  gopacket.SerializeBuffer
	options      gopacket.SerializeOptions
	pcktEthLayer layers.Ethernet
	pcktIPLayer  layers.IPv4
	pcktTCPLayer layers.TCP
)

// EncapsulatedPacket is used to store data in encapsulated packet
type EncapsulatedPacket struct {
	SrcIP  string          `json:"SrcIP"`
	DstIP  string          `json:"DstIP"`
	SrcMAC string          `json:"SrcMAC"`
	DstMAC string          `json:"DstMAC"`
	Packet gopacket.Packet `json:"Packet"`
}

// CreateAndSendEncapsulatedPacket is used to send a encapsulated packet
func CreateAndSendEncapsulatedPacket(handle *pcap.Handle, encapPacket EncapsulatedPacket) {
	ipLayer := &layers.IPv4{
		SrcIP: net.ParseIP(encapPacket.SrcIP),
		DstIP: net.ParseIP(encapPacket.DstIP),
	}
	srcMAC, _ := net.ParseMAC(encapPacket.SrcMAC)
	dstMAC, _ := net.ParseMAC(encapPacket.DstMAC)
	ethernetLayer := &layers.Ethernet{
		SrcMAC:       srcMAC,
		DstMAC:       dstMAC,
		EthernetType: layers.EthernetType(0x0800),
	}
	packetBuffer = gopacket.NewSerializeBuffer()
	err = gopacket.SerializePacket(packetBuffer, options, encapPacket.Packet)
	if err != nil {
		log.Fatal(err)
	}
	// And create the packet with the layers
	encapBuffer = gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(encapBuffer, options,
		ethernetLayer,
		ipLayer,
		&layers.TCP{},
		gopacket.Payload(packetBuffer.Bytes()),
	)
	outgoingPacket := encapBuffer.Bytes()
	err = handle.WritePacketData(outgoingPacket)
	if err != nil {
		log.Fatal(err)
	}
}

//GetPacketDstIP is used to get the destination ip
func GetPacketDstIP(packet gopacket.Packet) string {
	parser := gopacket.NewDecodingLayerParser(
		layers.LayerTypeEthernet,
		&pcktEthLayer,
		&pcktIPLayer,
		&pcktTCPLayer,
	)
	foundLayerTypes := []gopacket.LayerType{}

	err := parser.DecodeLayers(packet.Data(), &foundLayerTypes)
	if err != nil {
		fmt.Println("Trouble decoding layers: ", err)
	}

	for _, layerType := range foundLayerTypes {
		if layerType == layers.LayerTypeIPv4 {
			return string(pcktIPLayer.DstIP)
		}
	}
	return "Not Found"
}

// func tmp() {
// 	// Open device
// 	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer handle.Close()

// 	// Set filter
// 	var filter string = "icmp"
// 	err = handle.SetBPFFilter(filter)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Only capturing ICMP packets")
// 	// ip layer for forwarding packet
// 	ipLayer := &layers.IPv4{
// 		SrcIP: net.IP{192, 168, 0, 3},
// 		DstIP: net.IP{192, 168, 0, 4},
// 	}
// 	ethernetLayer := &layers.Ethernet{
// 		SrcMAC:       net.HardwareAddr{0xB8, 0x27, 0xEB, 0x81, 0x12, 0xEB},
// 		DstMAC:       net.HardwareAddr{0xB8, 0x27, 0xEB, 0x9A, 0x5E, 0xA5},
// 		EthernetType: layers.EthernetType(0x0800),
// 	}
// 	//rawBytes := []byte{10, 20, 30}

// 	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
// 	for packet := range packetSource.Packets() {
// 		fmt.Println("recieve")
// 		newBuffer := gopacket.NewSerializeBuffer()
// 		err := gopacket.SerializePacket(newBuffer, options, packet)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// And create the packet with the layers
// 		buffer = gopacket.NewSerializeBuffer()
// 		gopacket.SerializeLayers(buffer, options,
// 			ethernetLayer,
// 			ipLayer,
// 			&layers.TCP{},
// 			gopacket.Payload(newBuffer.Bytes()),
// 		)
// 		outgoingPacket := buffer.Bytes()
// 		err = handle.WritePacketData(outgoingPacket)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }
