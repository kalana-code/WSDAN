package flowmanager

import (
	"Agent/database"
)

// PacketDetails is used to store data of incomming packets
type PacketDetails struct {
	DstIP    string `json:"DstIP"`
	Protocol string `json:"Protocol"`
}

var defaultRule = database.RuleConfiguration{
	DstIP:    "any",
	Protocol: "any",
	FlowID:   "default",
	DstMAC:   "b8:27:eb:9a:5e:a5",
}

// RuleChecker is used to check the availablity of a rule
func RuleChecker(db map[string]database.RuleConfiguration, packetDetails PacketDetails) database.RuleConfiguration {
	for _, rule := range db {
		if rule.DstIP == packetDetails.DstIP && rule.Protocol == packetDetails.Protocol {
			return rule
		}
	}
	return defaultRule
}
