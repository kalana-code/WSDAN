package db

import (
	"Beq/rules/model"
	"errors"
	"fmt"
	"strconv"
	"sync"
)

//RuleDB is store rules
type RuleDB map[string]model.RulesDataRow

var instance RuleDB
var ruleIndex int
var flowID int

var once sync.Once

//GetRuleStore Initiating rule database
func GetRuleStore() *RuleDB {
	once.Do(func() {
		instance = make(map[string]model.RulesDataRow)
		ruleIndex = 0
	})
	return &instance
}

//AddRule add user to data base
func (obj *RuleDB) AddRule(User model.RulesDataRow) error {
	if instance != nil {
		instance[obj.getRuleID()] = User
		ruleIndex++
		return nil
	}
	return errors.New("No Data Base Initiate")

}

func (*RuleDB) getRuleID() string {
	return "R" + strconv.Itoa(ruleIndex)
}

func (*RuleDB) getFlowID() string {
	return "F" + strconv.Itoa(ruleIndex)
}

//FindRuleByRuleID used for get Rule bu RuleID
func (*RuleDB) FindRuleByRuleID(RuleID string) (*model.RulesDataRow, error) {
	if instance != nil {
		rule, pst := instance[RuleID]
		if !pst {
			return nil, errors.New("Rule NOT FOUND! ")
		}
		return &rule, nil
	}
	return nil, errors.New("No Data Base Initiate")
}

//RemoveRuleByRuleID used for remove Rule by RuleID
func (*RuleDB) RemoveRuleByRuleID(RuleID string) (string, *string, error) {
	if instance != nil {
		_, ok := instance[RuleID]
		if ok {
			NodeIP := instance[RuleID].NodeIP
			delete(instance, RuleID)
			return "Successfully Removed A Rule ", &NodeIP, nil
		}
		return "Not Exist Any Rule For Given RuleID ", nil, nil
	}
	return "No Data Base Initiate", nil, errors.New("No Data Base Initiate")
}

//RemoveRulesByFlowID used for remove Rule by RuleID
func (*RuleDB) RemoveRulesByFlowID(FlowID string) (string, error) {
	if instance != nil {
		isRemoved := false
		for RuleID, RuleData := range instance {
			if RuleData.FlowID == FlowID {
				delete(instance, RuleID)
				isRemoved = true
			}
		}
		if isRemoved {
			fmt.Print("S")
			return "Successfully Rules Belongs To Given FlowID", nil
		}
		return "Not Exist Any Flow For Given FlowID ", nil
	}
	return "", errors.New("No Data Base Initiate")
}

//GetAllRules used for remove Rule by RuleID
func (*RuleDB) GetAllRules() (*[]model.Rule, error) {
	rules := []model.Rule{}
	if instance != nil {
		for RuleID, RuleData := range instance {
			temp := model.Rule{}
			temp.Populate(RuleID, RuleData)
			rules = append(rules, temp)
		}
		return &rules, nil
	}
	return nil, errors.New("No Data Base Initiate")
}

//IsSet used for get Rule set State
func (*RuleDB) IsSet(RuleID string) (bool, error) {
	if instance != nil {
		rule, ok := instance[RuleID]
		if ok {

			return rule.IsSet, nil
		}
		return false, errors.New("Not Exist Any Rule For Given RuleID")
	}
	return false, errors.New("No Data Base Initiate")
}

//DispursedRule used for set Rule set State
func (*RuleDB) DispursedRule(RuleID string) error {
	if instance != nil {
		rule, ok := instance[RuleID]
		if ok {
			rule.IsSet = true
			return nil
		}
		return errors.New("Not Exist Any Rule For Given RuleID")
	}
	return errors.New("No Data Base Initiate")
}

//FindRuleByFlowID used for get Rule bu RuleID
func (*RuleDB) FindRuleByFlowID(FlowID string) (*[]model.RulesDataRow, error) {
	if instance != nil {
		rules := []model.RulesDataRow{}
		for _, RuleData := range instance {
			if RuleData.FlowID == FlowID {
				rules = append(rules, RuleData)
			}
		}
		if len(rules) == 0 {
			return nil, errors.New("User NOT FOUND! ")
		}
		return &rules, nil

	}
	return nil, errors.New("No Data Base Initiate")
}