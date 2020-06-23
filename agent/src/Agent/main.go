package main

import (
	"Agent/client"
	"Agent/config"
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
	device        string = "wlan0"
	snapshotLen   int32  = 1024
	promiscuous   bool   = false
	err           error
	handle        *pcap.Handle
	timeout       time.Duration = 1 * time.Second
	buffer        gopacket.SerializeBuffer
	infoLog       string = "INFO: [MN]:"
	errorLog      string = "ERROR: [MN]:"
	nodeName      string
	nodeGroup     string
	controllerMAC string
)

func main() {
	log.Println(infoLog, "------ Starting Node Functions ------")
	fmt.Println("Enter Node Name: ")
	fmt.Scanln(&nodeName)
	fmt.Println("Enter Node Group: ")
	fmt.Scanln(&nodeGroup)
	err = initializer.IptableInitializer()
	if err != nil {
		log.Println(errorLog, "Error when initializing iptables:", err)
	}
	controllerMAC, err = client.GetControllerMAC()
	if err != nil {
		log.Println(errorLog, "Error when retrieving controller MAC:", err)
		os.Exit(1)
	}
	settings := config.GetSettings()
	log.Println(infoLog, "Controller MAC : ", controllerMAC)
	err = settings.SetControllerMAC(controllerMAC)
	if err != nil {
		log.Println(errorLog, "Error when setting controller MAC", err)
	}
	client.SendNodeData(nodeName, nodeGroup)
	go doEvery(func() { client.SendNodeData(nodeName, nodeGroup) }, 300000*time.Millisecond)
	go server.Server()
	nfq, err := netfilter.NewNFQueue(0, 100, netfilter.NF_DEFAULT_PACKET_SIZE)
	if err != nil {
		log.Println(errorLog, "Error when initializing NFQueue:", err)
		os.Exit(1)
	}
	defer nfq.Close()
	packets := nfq.GetPackets()
	for {
		select {
		case p := <-packets:
			packet := p.Packet
			handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
			if err != nil {
				log.Println(errorLog, "OpenLive Error:", err)
			}
			defer handle.Close()
			buffer = packethandler.PacketAnalyzer(packet)
			p.SetVerdict(netfilter.NF_DROP)
			if buffer != nil {
				fmt.Println(gopacket.NewPacket(buffer.Bytes(), layers.LayerTypeEthernet, gopacket.Default))
				log.Println(infoLog, "Packet Sending")
				err = handle.WritePacketData(buffer.Bytes())
				if err != nil {
					log.Println(errorLog, "Packet Writing Error:", err)
				}
			} else {
				log.Println(infoLog, "Packet is Dropped")
			}
		}
	}
}

func doEvery(f func(), d time.Duration) {
	log.Println(infoLog, "Invoke doEvery")
	for range time.Tick(d) {
		f()
	}
}
