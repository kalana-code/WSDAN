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
	device       string = "wlan0"
	snapshot_len int32  = 1024
	promiscuous  bool   = false
	err          error
	timeout      time.Duration = 30 * time.Second
	handle       *pcap.Handle
	buffer       gopacket.SerializeBuffer
	options      gopacket.SerializeOptions
)

func main() {
	// Open device
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	fmt.Println("Processing!")
	//packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	sum := 1
	for sum < 1000 {
		sum += 1
		time.Sleep(2 * time.Second)
		//
		// Send raw bytes over wire
		rawBytes := []byte{10, 20, 30}

		// This time lets fill out some information
		ipLayer := &layers.IPv4{
			SrcIP: net.IP{192, 168, 0, 3},
			DstIP: net.IP{192, 168, 0, 4},
		}
		ethernetLayer := &layers.Ethernet{
			SrcMAC:       net.HardwareAddr{0xB8, 0x27, 0xEB, 0x81, 0x12, 0xEB},
			DstMAC:       net.HardwareAddr{0xB8, 0x27, 0xEB, 0x9A, 0x5E, 0xA5},
			EthernetType: layers.EthernetType(0x0800),
		}
		tcpLayer := &layers.TCP{
			SrcPort: layers.TCPPort(4321),
			DstPort: layers.TCPPort(8978),
		}
		// And create the packet with the layers
		buffer = gopacket.NewSerializeBuffer()
		gopacket.SerializeLayers(buffer, options,
			ethernetLayer,
			ipLayer,
			tcpLayer,
			gopacket.Payload(rawBytes),
		)
		outgoingPacket := buffer.Bytes()
		// Send our packet
		err = handle.WritePacketData(outgoingPacket)
		if err != nil {
			log.Fatal(err)
		}
	}
}
