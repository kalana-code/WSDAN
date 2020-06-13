package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/AkihiroSuda/go-netfilter-queue"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	device      string = "wlan0"
	snapshotLen int32  = 1024
	promiscuous bool   = false
	err         error
	handle      *pcap.Handle
	timeout     time.Duration = 1 * time.Second
	buffer      gopacket.SerializeBuffer
	options     gopacket.SerializeOptions
	// ethl *layers.Ethernet = &layers.Ethernet{}
	ipl         *layers.IPv4   = &layers.IPv4{}
	tcpl        *layers.ICMPv4 = &layers.ICMPv4{}
	dataPayload []byte         = nil
)

func main() {
	var err error

	nfq, err := netfilter.NewNFQueue(0, 100, netfilter.NF_DEFAULT_PACKET_SIZE)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer nfq.Close()
	packets := nfq.GetPackets()

	for true {
		select {
		case p := <-packets:
			fmt.Println(p.Packet)
			packet := p.Packet
			handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
			if err != nil {
				log.Fatal(err)
			}
			defer handle.Close()
			// ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
			// if ethernetLayer != nil {
			//     fmt.Println("ETH")
			//     ethl, _ = ethernetLayer.(*layers.Ethernet)
			// }
			// ethl.SrcMAC=net.HardwareAddr{0xB8, 0x27, 0xEB, 0x81, 0x12, 0xEB}
			// ethl.DstMAC=net.HardwareAddr{0xB8, 0x27, 0xEB, 0x9A, 0x5E, 0xA5}
			ethernetLayer := &layers.Ethernet{
				SrcMAC:       net.HardwareAddr{0xB8, 0x27, 0xEB, 0x81, 0x12, 0xEB},
				DstMAC:       net.HardwareAddr{0xB8, 0x27, 0xEB, 0x9A, 0x5E, 0xA5},
				EthernetType: layers.EthernetType(0x0800),
			}
			ipLayer := packet.Layer(layers.LayerTypeIPv4)
			if ipLayer != nil {
				fmt.Println("IPv4 layer detected.")
				ipl, _ = ipLayer.(*layers.IPv4)
			}

			tcpLayer := packet.Layer(layers.LayerTypeICMPv4)
			if tcpLayer != nil {
				fmt.Println("ICMP layer detected.")
				tcpl, _ = tcpLayer.(*layers.ICMPv4)
			}

			applicationLayer := packet.ApplicationLayer()
			if applicationLayer != nil {

				// Search for a string inside the payload
				if strings.Contains(string(applicationLayer.Payload()), "HTTP") {
					fmt.Println("HTTP found!")
				}
				dataPayload = applicationLayer.Payload()
			}
			options := gopacket.SerializeOptions{
				ComputeChecksums: true,
				FixLengths:       true,
			}

			buffer = gopacket.NewSerializeBuffer()
			//	tcpl.SetNetworkLayerForChecksum(ipl)
			err = gopacket.SerializeLayers(buffer, options,
				ethernetLayer,
				ipl,
				tcpl,
				gopacket.Payload(dataPayload),
			)
			// fmt.Println(er)
			if err != nil {
				panic(err)
			}
			//    fmt.Println(buffer.Bytes())
			//     fmt.Println("[184,39,235,154,94,165,184,39,235,129,18,235,8,0,0,0,0,0,0,0,0,0,0,0,0,0,192,168,0,3,192,168,0,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,184,39,235,129,18,235,184,39,235,83,110,202,8,0,69,0,0,84,198,90,64,0,64,1,242,249,192,168,0,1,192,168,0,3,8,0,120,221,13,208,16,46,162,163,216,94,241,30,10,0,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55]")
			// fmt.Println( gopacket.NewPacket( buffer.Bytes(), layers.LayerTypeEthernet, gopacket.Default))
			p.Packet = gopacket.NewPacket(buffer.Bytes(), layers.LayerTypeEthernet, gopacket.Default)
			fmt.Println(p.Packet)
			p.SetVerdict(netfilter.NF_DROP)
			err = handle.WritePacketData(buffer.Bytes())
			if err != nil {
				log.Fatal(err)
			}
			// p.SetVerdictWithPacket(netfilter.NF_ACCEPT,p.Packet.Bytes())
		}
	}
}
