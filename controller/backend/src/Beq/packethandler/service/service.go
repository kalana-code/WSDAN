package service

import (
	dispurseDB "Beq/dispurser/db"
	dispurseModel "Beq/dispurser/model"
	"Beq/packethandler/model"
	"Beq/packethandler/utils"
	ruleDB "Beq/rules/db"
	ruleModel "Beq/rules/model"
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
	infoLog     string = "INFO: [PS]:"
	errorLog    string = "ERROR: [PS]:"
)

// PacketAnalyzer is used to anaylize an input packet and create a new output packet
func PacketAnalyzer(packet gopacket.Packet) (gopacket.SerializeBuffer, *ruleModel.RulesDataRow) {
	log.Println(infoLog, "Starting PacketAnalyzer")
	packetDetails := model.PacketDetails{}
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		log.Println(infoLog, "IPv4 layer detected.")
		ipl, _ = ipLayer.(*layers.IPv4)
		packetDetails.SrcIP = ipl.SrcIP.String()
		packetDetails.DstIP = ipl.DstIP.String()
		packetDetails.Protocol = ipl.Protocol.String()
		log.Println(infoLog, "DstIP : ", packetDetails.DstIP, " SrcIP : ", ipl.SrcIP.String())
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
	ethernetLayer, rule := generateEthernetLayer(packetDetails)
	if ethernetLayer == nil {
		return nil, nil
	}
	if rule.Action == "ACCEPT" {
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
		return buffer, rule
	}
	return nil, rule
}

func generateEthernetLayer(packetDetails model.PacketDetails) (*layers.Ethernet, *ruleModel.RulesDataRow) {
	log.Println(infoLog, "Generating Ethernet Layer")
	rulesDb := ruleDB.GetRuleStore()
	ruleConfiguration, err := rulesDb.FindRuleByDstIPAndSrcIPAndProtocol(packetDetails)
	if err != nil {
		log.Println(errorLog, "DB data retrieving error", err)
		return nil, nil
	}
	if ruleConfiguration != nil {
		_, hardwareAddrs, _ := utils.GetIPAndMAC()
		srcMAC, _ := net.ParseMAC(hardwareAddrs)
		dstMAC, _ := net.ParseMAC(ruleConfiguration.DstMAC)
		ethernetLayer := &layers.Ethernet{
			SrcMAC:       srcMAC,
			DstMAC:       dstMAC,
			EthernetType: layers.EthernetType(0x0800),
		}
		return ethernetLayer, ruleConfiguration
	}
	return nil, nil
}

// DispurseFlow is used to dispurse rules to all other nodes
func DispurseFlow(flowID string) {
	log.Println(infoLog, "Dispurse rules corresponding to a flow")
	rulesDb := ruleDB.GetRuleStore()
	requestQueue := dispurseDB.GetRequestQueue()
	flowArray, err := rulesDb.FindRulesByFlowID(flowID)
	if err != nil {
		log.Println(errorLog, "DB data retrieving error", err)
	}
	for _, flow := range *flowArray {
		taskDetails := dispurseModel.AddRuleJob{
			RuleID:    flow.RuleID,
			Protocol:  flow.Protocol,
			FlowID:    flow.FlowID,
			SrcIP:     flow.SrcIP,
			DstIP:     flow.DstIP,
			Interface: flow.Interface,
			DstMAC:    flow.DstMAC,
			Action:    flow.Action,
			IsActive:  flow.IsActive,
		}
		dispureFlow := dispurseModel.Job{
			Type:        dispurseModel.TypeAddRule,
			NodeIP:      flow.NodeIP,
			TaskDetails: taskDetails,
		}
		requestQueue.AddJob(dispureFlow)
		rulesDb.DispursedRule(flow.RuleID)
	}
}
