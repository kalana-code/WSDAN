package database

import "fmt"

// RuleConfiguration is used to store data in the node database
type RuleConfiguration struct {
	EnacapPcktDstIP  string `json:"DstIP"`
	EnacapPcktDstMAC string `json:"DstMAC"`
	Action           string `json:"Action"`
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

// CheckRule is used to check the availablity of a rule
func CheckRule(db map[string]RuleConfiguration, key string) (*RuleConfiguration, bool) {
	if rule, found := db[key]; found {
		return &rule, found
	}
	return nil, false
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
