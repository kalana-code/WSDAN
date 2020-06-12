package database

import (
	"fmt"
	"log"
	"sync"
)

var (
	once     sync.Once
	instance map[string]RuleConfiguration
	infoLog  string = "INFO: [DB]:"
	errorLog string = "ERROR: [DB]:"
)

// RuleConfiguration is used to store data in the node database
type RuleConfiguration struct {
	DstIP    string `json:"DstIP"`
	Protocol string `json:"Protocol"`
	FlowID   string `json:"FlowID"`
	DstMAC   string `json:"DstMAC"`
}

// GetDatabase is used to get one instance of the db
func GetDatabase() map[string]RuleConfiguration {
	log.Println(infoLog, "Invoke GetDatabase")
	once.Do(func() {
		instance = make(map[string]RuleConfiguration)
	})
	return instance
}

// CreateRule is used to add a rule
func CreateRule(key string, newRule RuleConfiguration) {
	log.Println(infoLog, "Invoke CreateRule")
	db := GetDatabase()
	db[key] = newRule
}

// ViewRules is used to print the database
func ViewRules() {
	log.Println(infoLog, "Invoke ViewRules")
	db := GetDatabase()
	for key, value := range db {
		fmt.Println(key, value)
	}
}

// DeleteRule is used to delete a rule
func DeleteRule(key string) {
	log.Println(infoLog, "Invoke DeleteRule")
	db := GetDatabase()
	delete(db, key)
}
