package main

import (
	"fmt"
	"log"
	"net"
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
	timeout     time.Duration = 30 * time.Second
	handle      *pcap.Handle
	buffer      gopacket.SerializeBuffer
	options     gopacket.SerializeOptions
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
	fmt.Println("Only capturing ICMP packets")
	// ip layer for forwarding packet
	ipLayer := &layers.IPv4{
		SrcIP: net.IP{192, 168, 0, 3},
		DstIP: net.IP{192, 168, 0, 4},
	}
	ethernetLayer := &layers.Ethernet{
		SrcMAC:       net.HardwareAddr{0xB8, 0x27, 0xEB, 0x81, 0x12, 0xEB},
		DstMAC:       net.HardwareAddr{0xB8, 0x27, 0xEB, 0x9A, 0x5E, 0xA5},
		EthernetType: layers.EthernetType(0x0800),
	}
	//rawBytes := []byte{10, 20, 30}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println("recieve")
		newBuffer := gopacket.NewSerializeBuffer()
		err := gopacket.SerializePacket(newBuffer, options, packet)
		if err != nil {
			log.Fatal(err)
		}
		// And create the packet with the layers
		buffer = gopacket.NewSerializeBuffer()
		gopacket.SerializeLayers(buffer, options,
			ethernetLayer,
			ipLayer,
			&layers.TCP{},
			gopacket.Payload(newBuffer.Bytes()),
		)
		outgoingPacket := buffer.Bytes()
		err = handle.WritePacketData(outgoingPacket)
		if err != nil {
			log.Fatal(err)
		}
	}
}
