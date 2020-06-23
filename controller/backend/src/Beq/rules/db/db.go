package db

import (
	packethandlerModel "Beq/packethandler/model"
	"Beq/rules/model"
	"errors"
	"strconv"
	"sync"
)

//RuleDB is store rules
type RuleDB map[string]model.RulesDataRow

var instance RuleDB
var flowData map[string]model.FlowData
var ruleIndex int
var flowID int

var once sync.Once

//GetRuleStore Initiating rule database
func GetRuleStore() *RuleDB {
	once.Do(func() {
		instance = make(map[string]model.RulesDataRow)
		flowData = make(map[string]model.FlowData)
		ruleIndex = 0
	})
	return &instance
}

//AddRule add user to data base
func (obj *RuleDB) AddRule(rule model.RulesDataRow) (string, error) {
	if instance != nil {
		ruleID := obj.getRuleID()
		if flowData[rule.FlowID].DstIP == "" {
			flowData[rule.FlowID] = model.FlowData{
				DstIP: rule.DstIP,
				SrcIP: rule.SrcIP,
			}
		}
		instance[ruleID] = rule
		ruleIndex++
		return ruleID, nil
	}
	return "", errors.New("No Data Base Initiated")

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

//FindRuleByDstIPAndProtocol used for get Rule bu DstIP and Protocol
func (*RuleDB) FindRuleByDstIPAndProtocol(packetDetails packethandlerModel.PacketDetails) (*model.RulesDataRow, error) {
	if instance != nil {
		for _, rule := range instance {
			if rule.DstIP == packetDetails.DstIP && rule.Protocol == packetDetails.Protocol {
				return &rule, nil
			}
		}
		return nil, nil
	}
	return nil, errors.New("No Data Base Initiate")
}

//hasFlowID used to get rules by flow Id
func hasFlowID(flowIDVal string) bool {
	for _, RuleData := range instance {
		if RuleData.FlowID == flowIDVal {
			return true
		}
	}
	return false

}

//FindRulesByFlowID used to get rules by flow Id
func (*RuleDB) FindRulesByFlowID(flowIDVal string) (*[]model.Rule, error) {
	rules := []model.Rule{}
	if instance != nil {
		for RuleID, RuleData := range instance {
			if RuleData.FlowID == flowIDVal {
				temp := model.Rule{}
				temp.Populate(RuleID, RuleData)
				rules = append(rules, temp)
			}
		}
		return &rules, nil
	}
	return nil, errors.New("No Data Base Initiate")
}

//RemoveRuleByRuleID used for remove Rule by RuleID
func (*RuleDB) RemoveRuleByRuleID(RuleID string) (string, *string, error) {
	if instance != nil {
		_, ok := instance[RuleID]
		if ok {
			NodeIP := instance[RuleID].NodeIP
			FlowID := instance[RuleID].FlowID
			delete(instance, RuleID)
			isFlowExist := hasFlowID(FlowID)
			if !isFlowExist {
				delete(flowData, FlowID)
			}
			return "Successfully Removed A Rule ", &NodeIP, nil
		}
		return "Not Exist Any Rule For Given RuleID ", nil, nil
	}
	return "No Data Base Initiate", nil, errors.New("No Data Base Initiate")
}

//ChangeRuleStateByRuleID used for remove Rule by RuleID
func (*RuleDB) ChangeRuleStateByRuleID(RuleID string) (string, *string, bool, error) {
	if instance != nil {
		_, ok := instance[RuleID]
		if ok {
			rule := instance[RuleID]
			return "Successfully Removed A Rule ", &rule.NodeIP, rule.ChangeState(), nil
		}
		return "Not Exist Any Rule For Given RuleID ", nil, false, nil
	}
	return "No Data Base Initiate", nil, false, errors.New("No Data Base Initiate")
}

//RemoveRulesByFlowID used for remove Rule by RuleID
func (*RuleDB) RemoveRulesByFlowID(FlowID string) (string, error) {
	if instance != nil {
		isRemoved := false
		for RuleID, RuleData := range instance {
			if RuleData.FlowID == FlowID {
				delete(flowData, FlowID)
				delete(instance, RuleID)
				isRemoved = true
			}
		}
		if isRemoved {
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

//GetFlowData  used for geting flow data
func (*RuleDB) GetFlowData() map[string]model.FlowData {
	return flowData
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
