package database

import "fmt"

// RuleConfiguration is used to store data in the node database
type RuleConfiguration struct {
	DstIP    string `json:"DstIP"`
	Protocol string `json:"Protocol"`
	FlowID   string `json:"FlowID"`
	DstMAC   string `json:"DstMAC"`
}

// CreateDatabase is used to create a database
func CreateDatabase() map[string]RuleConfiguration {
	var database = make(map[string]RuleConfiguration)
	return database
}

// CreateRule is used to add a rule
func CreateRule(db map[string]RuleConfiguration, key string, newRule RuleConfiguration) {
	db[key] = newRule
}

// ViewRules is used to print the database
func ViewRules(db map[string]RuleConfiguration) {
	for key, value := range db {
		fmt.Println(key, value)
	}
}

// DeleteRule is used to delete a rule
func DeleteRule(db map[string]RuleConfiguration, key string) {
	delete(db, key)
}
