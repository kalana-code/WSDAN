package flowmanager

import (
	"Agent/config"
	"Agent/database"
	"log"
)

// PacketDetails is used to store data of incomming packets
type PacketDetails struct {
	DstIP    string `json:"DstIP"`
	SrcIP    string `json:"SrcIP"`
	Protocol string `json:"Protocol"`
}

// ControllerRuleConfiguration is used to store data in the node database
type ControllerRuleConfiguration struct {
	RuleID    string `json:"RuleID"`
	SrcIP     string `json:"SrcIP"`
	DstIP     string `json:"DstIP"`
	Protocol  string `json:"Protocol"`
	FlowID    string `json:"FlowID"`
	Interface string `json:"Interface"`
	DstMAC    string `json:"DstMAC"`
	Action    string `json:"Action"`
	IsActive  bool   `json:"IsActive"`
}

// RemoveRule is used to remove rule in the node database
type RemoveRule struct {
	RuleID string `json:"RuleID"`
}

// RuleState is used to set the isActive state of a rule
type RuleState struct {
	RuleID   string `json:"RuleID"`
	IsActive bool   `json:"IsActive"`
}

var (
	err           error
	controllerMAC string
	infoLog       string = "INFO: [FM]:"
	errorLog      string = "ERROR: [FM]:"
)

// RuleChecker is used to check the availablity of a rule
func RuleChecker(packetDetails PacketDetails) database.RuleConfiguration {
	log.Println(infoLog, "Invoke RuleChecker")
	db := database.GetDatabase()
	for _, rule := range db {
		if rule.DstIP == packetDetails.DstIP && rule.SrcIP == packetDetails.SrcIP &&
			rule.Protocol == packetDetails.Protocol && rule.IsActive {
			log.Println(infoLog, "A matching rule is found in DB")
			return rule
		}
	}
	log.Println(infoLog, "Default rule is set(Sending to the controller)")
	settings := config.GetSettings()
	controllerMAC, err = settings.GetControllerMAC()
	log.Println(infoLog, "Controller MAC", controllerMAC)
	defaultRule := database.RuleConfiguration{
		DstIP:     "any",
		SrcIP:     "any",
		Protocol:  "any",
		FlowID:    "default",
		Interface: "wlan0",
		DstMAC:    controllerMAC,
		Action:    "ACCEPT",
		IsActive:  true,
	}
	return defaultRule
}

// RuleUpdater is used to update a rule in database
func RuleUpdater(rule ControllerRuleConfiguration) {
	log.Println(infoLog, "Invoke RuleUpdater")
	newRuleConf := database.RuleConfiguration{
		DstIP:     rule.DstIP,
		SrcIP:     rule.SrcIP,
		Protocol:  rule.Protocol,
		FlowID:    rule.FlowID,
		Interface: rule.Interface,
		DstMAC:    rule.DstMAC,
		Action:    rule.Action,
		IsActive:  rule.IsActive,
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

// SetRuleState is used to set isActive state of a rule
func SetRuleState(ruleState RuleState) bool {
	log.Println(infoLog, "Invoke SetRuleState")
	isSet := database.SetRuleState(ruleState.RuleID, ruleState.IsActive)
	database.ViewRules()
	return isSet
}
