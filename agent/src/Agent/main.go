package main

import (
	"Agent/client"
	"Agent/database"
	"Agent/initializer"
	"Agent/packethandler"
	"Agent/server"
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
	infoLog     string = "INFO: [MN]:"
	errorLog    string = "ERROR: [MN]:"
)

func main() {
	log.Println(infoLog, "Starting Node Functions")
	log.Println("Test rule only .............................!!!")
	firstRule := database.RuleConfiguration{
		DstIP:    "192.168.0.2",
		Protocol: "ICMPv4",
		FlowID:   "124",
		DstMAC:   "b8:27:eb:5a:63:98",
	}
	database.CreateRule("1wesw", firstRule)
	err = initializer.IptableInitializer()
	if err != nil {
		log.Println(errorLog, "Error when initializing iptables:", err)
	}
	go doEvery(5000*time.Millisecond, client.SendNodeData)
	go server.Server()
	nfq, err := netfilter.NewNFQueue(0, 100, netfilter.NF_DEFAULT_PACKET_SIZE)
	if err != nil {
		log.Println(errorLog, "Error when initializing NFQueue:", err)
		os.Exit(1)
	}
	defer nfq.Close()
	packets := nfq.GetPackets()
	for{
		select {
		case p := <-packets:
			packet := p.Packet
			handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
			if err != nil {
				log.Println(errorLog, "OpenLive Error:", err)
			}
			defer handle.Close()
			buffer = packethandler.PacketAnalyzer(packet)
			fmt.Println(gopacket.NewPacket(buffer.Bytes(), layers.LayerTypeEthernet, gopacket.Default))
			p.SetVerdict(netfilter.NF_DROP)
			log.Println(infoLog, "Packet Sending")
			err = handle.WritePacketData(buffer.Bytes())
			if err != nil {
				log.Println(errorLog, "Packet Writing Error:", err)
			}
		}
	}
}

func doEvery(d time.Duration, f func()) {
	log.Println(infoLog, "Invoke doEvery")
	for range time.Tick(d) {
		f()
	}
}
