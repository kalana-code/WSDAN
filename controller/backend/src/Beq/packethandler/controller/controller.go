package controller

import (
	"Beq/packethandler/service"
	rulesModel "Beq/rules/model"
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
	rule		*rulesModel.RulesDataRow
	handle      *pcap.Handle
	timeout     time.Duration = 1 * time.Second
	buffer      gopacket.SerializeBuffer
	infoLog     string = "INFO: [PC]:"
	errorLog    string = "ERROR: [PC]:"
)

func packetController() {

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
			buffer, rule = service.PacketAnalyzer(packet)
			fmt.Println(gopacket.NewPacket(buffer.Bytes(), layers.LayerTypeEthernet, gopacket.Default))
			p.SetVerdict(netfilter.NF_DROP)
			if buffer != nil {
				if rule.IsSet {
					log.Println(infoLog, "Rule is already set in Nodes")
					}else{
						log.Println(infoLog, "send the rule and set setRule=true")
					}
				log.Println(infoLog, "Packet Sending")
				err = handle.WritePacketData(buffer.Bytes())
				if err != nil {
					log.Println(errorLog, "Packet Writing Error:", err)
				}
			}else{
				log.Println(infoLog, "Packet Dropped")
			}
		}
	}
}
}