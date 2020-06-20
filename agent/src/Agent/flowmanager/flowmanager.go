package flowmanager

import (
	"Agent/database"
	"log"
)

// PacketDetails is used to store data of incomming packets
type PacketDetails struct {
	DstIP    string `json:"DstIP"`
	Protocol string `json:"Protocol"`
}

// ControllerRuleConfiguration is used to store data in the node database
type ControllerRuleConfiguration struct {
	RuleID    string `json:"RuleID"`
	DstIP     string `json:"DstIP"`
	Protocol  string `json:"Protocol"`
	FlowID    string `json:"FlowID"`
	Interface string `json:"Interface"`
	DstMAC    string `json:"DstMAC"`
	Action    string `json:"Action"`
}

// RemoveRule is used to remove rule in the node database
type RemoveRule struct {
	RuleID string `json:"RuleID"`
}

var (
	defaultRule = database.RuleConfiguration{
		DstIP:     "any",
		Protocol:  "any",
		FlowID:    "default",
		Interface: "wlan0",
		DstMAC:    "b8:27:eb:9a:5e:a5",
		Action:    "ACCEPT",
	}
	infoLog  string = "INFO: [FM]:"
	errorLog string = "ERROR: [FM]:"
)

// RuleChecker is used to check the availablity of a rule
func RuleChecker(packetDetails PacketDetails) database.RuleConfiguration {
	log.Println(infoLog, "Invoke RuleChecker")
	db := database.GetDatabase()
	for _, rule := range db {
		if rule.DstIP == packetDetails.DstIP && rule.Protocol == packetDetails.Protocol {
			return rule
		}
	}
	return defaultRule
}

// RuleUpdater is used to update a rule in database
func RuleUpdater(rule ControllerRuleConfiguration) {
	log.Println(infoLog, "Invoke RuleUpdater")
	newRuleConf := database.RuleConfiguration{
		DstIP:     rule.DstIP,
		Protocol:  rule.Protocol,
		FlowID:    rule.FlowID,
		Interface: rule.Interface,
		DstMAC:    rule.DstMAC,
		Action:    rule.Action,
	}
	database.CreateRule(rule.RuleID, newRuleConf)
	database.ViewRules()
}

// RuleRemoveByRuleID is used to remove a rule by RuleID in database
func RuleRemoveByRuleID(removeRule RemoveRule) {
	log.Println(infoLog, "Invoke RuleRemoveByRuleID")
	database.DeleteRule(removeRule.RuleID)
	database.ViewRules()
}
