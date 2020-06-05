package main

import (
	"Agent/database"
	"Agent/packethandler"
	"fmt"
	"log"
	"os"
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
)

func main() {
	log.Println("Starting Node Functions ...!")
	db := database.CreateDatabase()
	firstRule := database.RuleConfiguration{
		DstIP:    "192.168.0.2",
		Protocol: "ICMPv4",
		FlowID:   "124",
		DstMAC:   "b8:27:eb:5a:63:98",
	}
	database.CreateRule(db, "1wesw", firstRule)
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
			buffer = packethandler.PacketAnalyzer(db, packet)
			fmt.Println(gopacket.NewPacket(buffer.Bytes(), layers.LayerTypeEthernet, gopacket.Default))
			p.SetVerdict(netfilter.NF_DROP)
			log.Println("Packet Sending ...!")
			err = handle.WritePacketData(buffer.Bytes())
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
